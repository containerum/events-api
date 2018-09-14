package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func (h *EventsHandlers) GetNamespaceChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetDeploymentChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetDeploymentChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetDeploymentChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceDeploymentsChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceDeploymentsChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceDeploymentsChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetServiceChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetServiceChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetServiceChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceServicesChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceServicesChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceServicesChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetIngressChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetIngressChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetIngressChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceIngressesChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceIngressesChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceIngressesChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetPVCChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetPVCChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetPVCChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespacePVCsChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespacePVCsChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespacePVCsChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetSecretChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetSecretChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetSecretChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceSecretsChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceSecretsChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceSecretsChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetConfigMapChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetConfigMapChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetConfigMapChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) GetNamespaceConfigMapsChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceConfigMapsChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			logrus.Warn(err)
		}
		resp, err := h.GetNamespaceConfigMapsChanges(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
