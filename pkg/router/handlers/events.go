package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/containerum/events-api/pkg/eaerrors"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin/binding"

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
		if err := withWS(ctx, h.GetPodEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetPodEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespacePodsEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespacePodsEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespacePodsEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetPVCEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetPVCEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetPVCEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespacePVCsEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespacePVCsEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespacePVCsEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
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

func (h *EventsHandlers) GetUsersEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetUsersEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetUsersEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetSystemEventsListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetSystemEvents); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetSystemEvents(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
