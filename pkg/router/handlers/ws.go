package handlers

import (
	"net/http"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/ninedraft/gocontrol"

	"github.com/containerum/events-api/pkg/util/ticker"
	"github.com/containerum/kube-client/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func withWS(ctx *gin.Context, limit int, startTime time.Time, getfuncs ...eventsFunc) {
	var control = &gocontrol.Guard{}
	defer control.Wait()
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Debug(err)
		return
	}
	defer c.Close()
	defer ctx.Abort()

	var errChan = make(chan error)
	var resultChan = make(chan model.Event)
	var finalChan = make(chan []model.Event)
	defer close(resultChan)
	defer close(errChan)
	defer close(finalChan)

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
		go EventAggregator{
			Ctx:         ctx,
			Limit:       limit,
			EventSource: eventSource,
			EventDrain:  resultChan,
			ErrChan:     errChan,
			Control:     control,
			StartAt:     startTime,
		}.Run()
	}

	go EventBatcher{
		Ctx:               ctx,
		ErrChan:           errChan,
		Quant:             1 * time.Second,
		BatchDrain:        finalChan,
		PreallocBatchSize: 16,
		EventSource:       resultChan,
		Control:           control,
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
						return timei.Before(timej)
					})
					limitOnce.Do(func() {
						if limit < len(result) {
							result = result[:limit]
						}
					})

					logrus.Infof("Writing %v events", len(result))
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
			logrus.Debug(selecerr)
			return
		}
	}
}

type EventBatcher struct {
	Ctx               *gin.Context
	ErrChan           chan error
	Quant             time.Duration
	EventSource       <-chan model.Event
	BatchDrain        chan<- []model.Event
	Control           *gocontrol.Guard
	PreallocBatchSize int
}

func (batcher EventBatcher) Run() {
	defer batcher.Control.Go()()

	var ctx = batcher.Ctx
	var aborted = AbortWaiter(ctx.IsAborted)
	var finalChan = batcher.BatchDrain
	var resultChan = batcher.EventSource

	results := make([]model.Event, 0)
	var timer = ticker.NewTicker(batcher.Quant)
	timer.Start()
	defer timer.Stop()
	for {
		select {
		case <-aborted:
			return
		case <-ctx.Done():
			return
		case <-timer.Ticks():
			finalChan <- results
			results = make([]model.Event, 0)
		case event, ok := <-resultChan:
			if !ok {
				return
			}
			results = append(results, event)
		}
	}
}

type EventAggregator struct {
	Ctx         *gin.Context
	StartAt     time.Time
	Limit       int
	EventSource eventsFunc
	EventDrain  chan<- model.Event
	ErrChan     chan<- error
	Control     *gocontrol.Guard
}

func (aggregate EventAggregator) Run() {
	defer aggregate.Control.Go()()

	var ctx = aggregate.Ctx
	var getfunc = aggregate.EventSource
	var funcLimit = aggregate.Limit
	var startTime = aggregate.StartAt
	var resultChan = aggregate.EventDrain
	var errChan = aggregate.ErrChan
	var aborted = AbortWaiter(aggregate.Ctx.IsAborted)
	var firstEventSend = false

	for {
		select {
		case <-aborted: //If context aborted finish goroutine
			return
		case <-aggregate.Ctx.Done():
			return
		default:
			resp, err := getfunc(ctx.Params, funcLimit, startTime)
			if err != nil {
				errChan <- err
				return
			}
		batchLoop:
			for _, event := range resp.Events {
				select {
				case <-aborted: //If context aborted finish goroutine
					return
				case <-aggregate.Ctx.Done():
					return
				case resultChan <- event:
					continue batchLoop
				}
			}
			if !firstEventSend {
				firstEventSend = true
			}
			funcLimit = 0
			//Get only new events
			startTime = time.Now()
			time.Sleep(60 * time.Second)
		}
	}
}

func AbortWaiter(aborted func() bool) <-chan struct{} {
	var wait = make(chan struct{})
	go func() {
		defer close(wait)
		for {
			if aborted() {
				return
			}
			runtime.Gosched()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return wait
}
