package handlers

import (
	"sort"
	"time"

	"github.com/containerum/kube-client/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

func withWS(ctx *gin.Context, limit int, getfuncs ...eventsFunc) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Debug(err)
		return
	}
	defer c.Close()
	//Abort context in the end
	defer ctx.Abort()

	var errChan = make(chan error)
	var resultChan = make(chan model.Event)
	var finalChan = make(chan []model.Event)

	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil
	})

	//Checking for closed connection
	go func() {
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	//Limiter. Waiting for all DB request to finish on first run.
	firstRun := len(getfuncs)

	for _, getfunc := range getfuncs {
		//Get events from the beginning for the first time with default limit
		var funcLimit = limit
		var startTime time.Time
		go func(getfunc eventsFunc) {
			for {
				//If context aborted finish goroutine
				if ctx.IsAborted() {
					return
				}
				resp, err := getfunc(ctx.Params, funcLimit, startTime)
				if firstRun > 0 {
					firstRun--
				}
				if err != nil {
					errChan <- err
					return
				}
				//Check again
				if ctx.IsAborted() {
					return
				}
				for _, event := range resp.Events {
					resultChan <- event
				}
				funcLimit = 0
				//Get only new events
				startTime = time.Now()
				time.Sleep(60 * time.Second)
			}
		}(getfunc)
	}

	go func() {
		results := make([]model.Event, 0)
		timer := time.NewTicker(100 * time.Millisecond)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				if firstRun > 0 {
					continue
				}
				timer = time.NewTicker(10 * time.Millisecond)
				finalChan <- results
				results = make([]model.Event, 0)
			case event, ok := <-resultChan:
				if !ok {
					close(finalChan)
					return
				}
				results = append(results, event)
			}
		}
	}()

	go func() {
		pingTicker := time.NewTicker(2 * time.Second)
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
					logrus.Infof("Writing %v events", len(result))
					if err := c.WriteJSON(result); err != nil {
						errChan <- err
					}
				}
			case <-pingTicker.C:
				if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
			}
		}
	}()

	for {
		select {
		case selecerr := <-errChan:
			if selecerr != nil {
				logrus.Debug(selecerr)
				return
			}
		}
	}
}
