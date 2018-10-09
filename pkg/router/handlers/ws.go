package handlers

import (
	"context"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/containerum/events-api/pkg/util/atomicbool"

	"github.com/containerum/events-api/pkg/model"

	"github.com/ninedraft/gocontrol"

	"github.com/containerum/events-api/pkg/util/ticker"
	kubeModel "github.com/containerum/kube-client/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func withWS(ginCtx *gin.Context, params model.FuncParams, dbPeriod time.Duration, getfuncs ...model.EventsFunc) {
	var control = &gocontrol.Guard{}
	defer control.Wait()
	c, err := upgrader.Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if err != nil {
		log.Debug(err)
		return
	}
	defer c.Close()

	ctx, cancelCTX := context.WithCancel(ginCtx.Request.Context())
	defer cancelCTX()

	var errChan = make(chan error)
	var resultChan = make(chan kubeModel.Event)
	var finalChan = make(chan []kubeModel.Event)
	defer close(resultChan)
	defer close(errChan)
	defer close(finalChan)

	var firstTimeWG = &sync.WaitGroup{}

	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil
	})

	//Checking for closed connection
	go func() {
		defer control.Go()()
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	//Limiter. Waiting for all DB request to finish on first run.
	for _, eventSource := range getfuncs {
		firstTimeWG.Add(1)
		go EventAggregator{
			CTX:                ctx,
			Params:             params,
			EventSource:        eventSource,
			EventDrain:         resultChan,
			ErrChan:            errChan,
			Control:            control,
			DBPeriod:           dbPeriod,
			FirstTimeWaitGroup: firstTimeWG,
		}.Run()
	}

	go EventBatcher{
		Ctx:                ctx,
		ErrChan:            errChan,
		Quant:              1 * time.Second,
		BatchDrain:         finalChan,
		PreallocBatchSize:  16,
		EventSource:        resultChan,
		Control:            control,
		FirstTimeWaitGroup: firstTimeWG,
	}.Run()

	go func() {
		defer control.Go()()

		var limitOnce = sync.Once{}

		pingTicker := ticker.NewTicker(1 * time.Second)
		pingTicker.Start()
		defer pingTicker.Stop()
		for {
			select {
			case result, ok := <-finalChan:
				if !ok {
					return
				}
				if len(result) > 0 {
					sort.Slice(result, func(i, j int) bool {
						timei, _ := time.Parse(time.RFC3339, result[i].Time)
						timej, _ := time.Parse(time.RFC3339, result[j].Time)
						return timei.After(timej)
					})
					limitOnce.Do(func() {
						if params.Limit < len(result) && params.Limit > 0 {
							result = result[:params.Limit]
						}
					})

					log.Infof("Writing %v events", len(result))
					if err := c.WriteJSON(result); err != nil {
						errChan <- err
					}
				}
			case <-pingTicker.Ticks():
				if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
			}
		}
	}()

	for selecerr := range errChan {
		if selecerr != nil {
			log.Debug(selecerr)
			return
		}
	}
}

type EventBatcher struct {
	Ctx                context.Context
	ErrChan            chan error
	Quant              time.Duration
	EventSource        <-chan kubeModel.Event
	BatchDrain         chan<- []kubeModel.Event
	Control            *gocontrol.Guard
	PreallocBatchSize  int
	FirstTimeWaitGroup *sync.WaitGroup
}

func (batcher EventBatcher) Run() {
	defer batcher.Control.Go()()

	var ctx = batcher.Ctx
	var finalChan = batcher.BatchDrain
	var resultChan = batcher.EventSource

	ready, checkReady := atomicbool.Create()
	go func() {
		batcher.FirstTimeWaitGroup.Wait()
		ready()
	}()

	results := make([]kubeModel.Event, 0)
	var timer = ticker.NewTicker(batcher.Quant)
	timer.Start()
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.Ticks():
			//Wait for first batch of events
			if !checkReady() {
				continue
			}
			finalChan <- results
			results = make([]kubeModel.Event, 0)
		case event, ok := <-resultChan:
			if !ok {
				return
			}
			results = append(results, event)
		}
	}
}

type EventAggregator struct {
	CTX                context.Context
	Params             model.FuncParams
	EventSource        model.EventsFunc
	EventDrain         chan<- kubeModel.Event
	ErrChan            chan<- error
	Control            *gocontrol.Guard
	DBPeriod           time.Duration
	FirstTimeWaitGroup *sync.WaitGroup
}

func (aggregate EventAggregator) Run() {
	defer aggregate.Control.Go()()

	var getfunc = aggregate.EventSource
	var resultChan = aggregate.EventDrain
	var errChan = aggregate.ErrChan
	var params = aggregate.Params
	var dbPeriod = aggregate.DBPeriod
	var firstTimeReadyOnce = sync.Once{}

	for {
		select {
		case <-aggregate.CTX.Done():
			return
		default:
			resp, err := getfunc(params)
			//Indicate that we got first batch from DB
			firstTimeReadyOnce.Do(func() {
				aggregate.FirstTimeWaitGroup.Done()
			})
			if err != nil {
				errChan <- err
				return
			}
		batchLoop:
			for _, event := range resp.Events {
				select {
				case <-aggregate.CTX.Done():
					return
				case resultChan <- event:
					continue batchLoop
				}
			}
			params.Limit = 0
			//Get only new events
			params.StartTime = time.Now()
			time.Sleep(dbPeriod)
		}
	}
}
