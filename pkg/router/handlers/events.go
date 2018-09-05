package handlers

import (
	"net/http"

	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
}

func (h *EventsHandlers) GetDeploymentEventsListHandler(ctx *gin.Context) {
	resp, err := h.GetDeploymentEvents(ctx.Request.Context(), ctx.Param("namespace"), ctx.Param("deployment"))
	if err != nil {
		ctx.AbortWithStatusJSON(h.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *EventsHandlers) GetNamespaceDeploymentsEventsListHandler(ctx *gin.Context) {
	resp, err := h.GetNamespaceDeploymentsEvents(ctx.Request.Context(), ctx.Param("namespace"))
	if err != nil {
		ctx.AbortWithStatusJSON(h.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *EventsHandlers) GetPodEventsListHandler(ctx *gin.Context) {
	resp, err := h.GetPodEvents(ctx.Request.Context(), ctx.Param("namespace"), ctx.Param("pod"))
	if err != nil {
		ctx.AbortWithStatusJSON(h.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *EventsHandlers) GetNamespacePodsEventsListHandler(ctx *gin.Context) {
	resp, err := h.GetNamespacePodsEvents(ctx.Request.Context(), ctx.Param("namespace"))
	if err != nil {
		ctx.AbortWithStatusJSON(h.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
