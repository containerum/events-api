package handlers

import (
	"runtime"
	"sort"
	"time"

	"github.com/containerum/events-api/pkg/eventfilter"
	"github.com/containerum/events-api/pkg/util/ticker"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ninedraft/gocontrol"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

func withWS(ctx *gin.Context, limit int, getfuncs ...eventsFunc) {
	var control = &gocontrol.Guard{}
	defer control.Wait()

	//Divide limit to all functions
	limit = limit / len(getfuncs)

	var filter eventfilter.Predicate = eventfilter.True
	var userDefinedLevelWhiteList = ctx.QueryArray("levels")
	if len(userDefinedLevelWhiteList) > 0 {
		var levelWhiteList = func() []model.EventKind {
			var levels = ctx.QueryArray("levels")
			var kinds = make([]model.EventKind, 0, len(levels))
			for _, level := range levels {
				kinds = append(kinds, model.EventKind(level))
			}
			return kinds
		}()
		filter = eventfilter.MatchAnyKind(levelWhiteList...)
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Debug(err)
		return
	}
	defer conn.Close()
	defer ctx.Abort()

	var errChan = make(chan error)
	var unfilteredEvents = make(chan model.Event)
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
		go EventAgregator{
			Ctx:         ctx,
			Limit:       limit,
			EventSource: eventSource,
			EventDrain:  unfilteredEvents,
			ErrChan:     errChan,
			Control:     control,
		}.Run()
	}

	var filteredEvents = filter.Pipe(-1, unfilteredEvents)
	var eventBatches = make(chan []model.Event)
	go EventBatcher{
		Ctx:               ctx,
		ErrChan:           errChan,
		Quant:             1 * time.Second,
		BatchDrain:        eventBatches,
		PreallocBatchSize: 16,
		EventSource:       filteredEvents,
		Сontrol:           control,
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
						return timei.Before(timej)
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
	Ctx               *gin.Context
	ErrChan           chan error
	Quant             time.Duration
	EventSource       <-chan model.Event
	BatchDrain        chan<- []model.Event
	Сontrol           *gocontrol.Guard
	PreallocBatchSize int
}

func (batcher EventBatcher) Run() {
	defer batcher.Сontrol.Go()()

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

type EventAgregator struct {
	Ctx         *gin.Context
	StartAt     time.Time
	Limit       int
	EventSource eventsFunc
	EventDrain  chan<- model.Event
	ErrChan     chan<- error
	Control     *gocontrol.Guard
}

func (agregate EventAgregator) Run() {
	defer agregate.Control.Go()()

	var ctx = agregate.Ctx
	var getfunc = agregate.EventSource
	var funcLimit = agregate.Limit
	var startTime = agregate.StartAt
	var resultChan = agregate.EventDrain
	var errChan = agregate.ErrChan
	var aborted = AbortWaiter(agregate.Ctx.IsAborted)
	var firstEventSend = false

	for {
		select {
		case <-aborted: //If context aborted finish goroutine
			return
		case <-agregate.Ctx.Done():
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
				case <-agregate.Ctx.Done():
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
