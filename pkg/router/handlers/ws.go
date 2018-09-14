package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

func withWS(ctx *gin.Context, getfunc eventsFunc, limit int) error {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	var startTime time.Time
	for {
		resp, err := getfunc(ctx.Params, limit, startTime)
		limit = 0
		startTime = time.Now()
		if err != nil {
			return err
		}
		if len(resp.Events) > 0 {
			logrus.Infof("Writing %v events", len(resp.Events))
			if err := c.WriteJSON(resp.Events); err != nil {
				return err
			}
		}
		time.Sleep(30 * time.Second)
	}
}
