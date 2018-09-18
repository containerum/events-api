package handlers

import (
	"time"

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
				if err != nil {
					errChan <- err
				}
				//Now we don't need limit
				funcLimit = 0
				//Get only new events
				startTime = time.Now()
				if len(resp.Events) > 0 {
					logrus.Infof("Writing %v events", len(resp.Events))
					if err := c.WriteJSON(resp.Events); err != nil {
						errChan <- err
					}
				}
				time.Sleep(30 * time.Second)
			}
		}(getfunc)
	}

	for {
		selecerr := <-errChan
		if selecerr != nil {
			logrus.Debug(selecerr)
			return
		}
	}
}
