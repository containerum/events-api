package handlers

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type eventsFunc func(gin.Params, time.Time) (*model.EventsList, error)

func withWS(ctx *gin.Context, getfunc eventsFunc) error {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	var startTime time.Time
	for {
		resp, err := getfunc(ctx.Params, startTime)
		startTime = time.Now()
		if err != nil {
			return err
		}
		if len(resp.Events) > 0 {
			logrus.Infof("Writing %v events", len(resp.Events))
			c.WriteJSON(resp.Events)
		}
		time.Sleep(30 * time.Second)
	}
}
