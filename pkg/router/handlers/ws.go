package handlers

import (
	"context"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/containerum/events-api/pkg/eventfilter"
	"github.com/containerum/events-api/pkg/model"
	"github.com/containerum/events-api/pkg/util/atomicbool"
	"github.com/containerum/events-api/pkg/util/ticker"
	kubemodel "github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ninedraft/gocontrol"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func withWS(ginCtx *gin.Context, params model.FuncParams, dbPeriod time.Duration, getfuncs ...model.EventsFunc) {
	var control = &gocontrol.Guard{}
	defer control.Wait()

	//Divide limit to all functions

	var filter eventfilter.Predicate = eventfilter.True
	var userDefinedLevelWhiteList = ginCtx.QueryArray("levels")
	var userDefinedSourceWhiteList = ginCtx.QueryArray("resources")

	if len(userDefinedLevelWhiteList) > 0 || len(userDefinedSourceWhiteList) > 0 {
		var levelWhiteList = func() []kubemodel.EventKind {
			var kinds = make([]kubemodel.EventKind, 0, len(userDefinedLevelWhiteList))
			for _, level := range userDefinedLevelWhiteList {
				kinds = append(kinds, kubemodel.EventKind(level))
			}
			return kinds
		}()

		var resourceWhiteList = func() []kubemodel.ResourceType {
			var resources = make([]kubemodel.ResourceType, 0, len(userDefinedSourceWhiteList))
			for _, resource := range userDefinedSourceWhiteList {
				resources = append(resources, kubemodel.ResourceType(resource))
			}
			return resources
		}()

		if len(userDefinedLevelWhiteList) > 0 {
			filter = eventfilter.MatchAnyKind(levelWhiteList...)
		}
		if len(userDefinedSourceWhiteList) > 0 {
			filter = eventfilter.And(filter, eventfilter.MatchAnyType(resourceWhiteList...))
		}
	}

	conn, err := upgrader.Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if err != nil {
		logrus.Debug(err)
		return
	}
	defer conn.Close()
	var errChan = make(chan error)
	var unfilteredEvents = make(chan kubemodel.Event)
	defer close(unfilteredEvents)
	defer close(errChan)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil
	})

	//Checking for closed connection
	go CheckConnection(control, conn, errChan)

	//Limiter. Waiting for all DB request to finish on first run.
	for _, eventSource := range getfuncs {
		go EventAggregator{
			CTX:         ginCtx,
			EventSource: eventSource,
			EventDrain:  unfilteredEvents,
			ErrChan:     errChan,
			Control:     control,
		}.Run()
	}

	var filteredEvents = filter.Pipe(-1, unfilteredEvents)
	var eventBatches = make(chan []kubemodel.Event)
	go EventBatcher{
		Ctx:               ginCtx,
		ErrChan:           errChan,
		Quant:             1 * time.Second,
		BatchDrain:        eventBatches,
		PreallocBatchSize: 16,
		EventSource:       filteredEvents,
	}.Run()

	go func() {
		defer control.Go()()

		pingTicker := ticker.NewTicker(1 * time.Second)
		pingTicker.Start()
		defer pingTicker.Stop()
		for {
			select {
			case result, ok := <-eventBatches:
				if !ok {
					return
				}
				if len(result) > 0 {
					sort.Slice(result, func(i, j int) bool {
						timei, _ := time.Parse(time.RFC3339, result[i].Time)
						timej, _ := time.Parse(time.RFC3339, result[j].Time)
						return timei.After(timej)
					})
					logrus.Infof("Writing %v events", len(result))
					if err := conn.WriteJSON(result); err != nil {
						errChan <- err
					}
				}
			case <-pingTicker.Ticks():
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
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
	Ctx                context.Context
	ErrChan            chan error
	Quant              time.Duration
	EventSource        <-chan kubemodel.Event
	BatchDrain         chan<- []kubemodel.Event
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

	results := make([]kubemodel.Event, 0)
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
			results = make([]kubemodel.Event, 0)
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
	EventDrain         chan<- kubemodel.Event
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

func CheckConnection(control *gocontrol.Guard, conn *websocket.Conn, errChan chan<- error) {
	defer control.Go()()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			errChan <- err
			return
		}
	}
}
