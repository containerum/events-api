package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *EventsHandlers) GetNamespaceChangesListHandler(ctx *gin.Context) {
	if _, ws := ctx.GetQuery("ws"); ws {
		if err := withWS(ctx, h.GetNamespaceChanges); err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
	} else {
		resp, err := h.GetNamespaceChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetDeploymentChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetNamespaceDeploymentsChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetServiceChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetNamespaceServicesChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetIngressChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetNamespaceIngressesChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetPVCChanges(ctx.Params, time.Time{})
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
		resp, err := h.GetNamespacePVCsChanges(ctx.Params, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
