package handlers

import (
	"net/http"
	"strconv"
	"time"

	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
}

type eventsFunc func(gin.Params, int, time.Time) (*model.EventsList, error)

func handleResourceChangesEvents(h *EventsHandlers, ctx *gin.Context, getFunc eventsFunc) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, getFunc, limit); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		resp, err := getFunc(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
