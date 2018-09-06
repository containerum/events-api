package handlers

import (
	"net/http"
	"time"

	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
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
		withWS(ctx, h.GetNamespacePodsEvents)
	} else {
		resp, err := h.GetNamespacePodsEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetPVCEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetPVCEvents)
	} else {
		resp, err := h.GetPVCEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespacePVCsEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, h.GetNamespacePVCsEvents)
	} else {
		resp, err := h.GetNamespacePVCsEvents(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
