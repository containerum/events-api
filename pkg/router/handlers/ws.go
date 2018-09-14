package handlers

import (
	"strconv"
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

type eventsFunc func(gin.Params, int, time.Time) (*model.EventsList, error)

func withWS(ctx *gin.Context, getfunc eventsFunc) error {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
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
