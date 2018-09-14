package handlers

import (
	"net/http"

	"github.com/containerum/events-api/pkg/eaerrors"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func (h *EventsHandlers) GetPodEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPodEvents)
}

func (h *EventsHandlers) GetNamespacePodsEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePodsEvents)
}

func (h *EventsHandlers) GetPVCEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPVCEvents)
}

func (h *EventsHandlers) GetNamespacePVCsEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePVCsEvents)
}

func (h *EventsHandlers) GetUsersEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetUsersEvents)
}

func (h *EventsHandlers) GetSystemEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetSystemEvents)
}

func (h *EventsHandlers) AddUserEventHandler(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindWith(&event, binding.JSON); err != nil {
		ctx.Error(err)
		gonic.Gonic(eaerrors.ErrValidation(), ctx)
		return
	}
	if err := h.AddUserEvent(event); err != nil {
		ctx.Error(err)
		gonic.Gonic(eaerrors.ErrUnableAddEvent(), ctx)
		return
	}
	ctx.Status(http.StatusAccepted)
}

func (h *EventsHandlers) AddSystemEventHandler(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindWith(&event, binding.JSON); err != nil {
		ctx.Error(err)
		gonic.Gonic(eaerrors.ErrValidation(), ctx)
		return
	}
	if err := h.AddSystemEvent(event); err != nil {
		ctx.Error(err)
		gonic.Gonic(eaerrors.ErrUnableAddEvent(), ctx)
		return
	}
	ctx.Status(http.StatusAccepted)
}
