package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/containerum/kube-client/pkg/model"

	"github.com/gorilla/websocket"

	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
}

type eventsFunc func(gin.Params, time.Time) (*model.EventsList, error)

func (h *EventsHandlers) GetDeploymentEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetDeploymentEvents)
	} else {
		resp, err := h.GetDeploymentEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceDeploymentsEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetNamespaceDeploymentsEvents)
	} else {
		resp, err := h.GetNamespaceDeploymentsEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetPodEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetPodEvents)
	} else {
		resp, err := h.GetPodEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespacePodsEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetPodEvents)
	} else {
		resp, err := h.GetNamespacePodsEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

var upgrader = websocket.Upgrader{}

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
			fmt.Printf("Writing %v events\n", len(resp.Events))
			c.WriteJSON(resp.Events)
		}
		time.Sleep(30 * time.Second)
	}
}
