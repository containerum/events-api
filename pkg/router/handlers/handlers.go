package handlers

import (
	"net/http"
	"strconv"
	"strings"
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

type eventsFunc func(params gin.Params, limit int, startFrom time.Time) (*model.EventsList, error)

func handleResourceChangesEvents(h *EventsHandlers, ctx *gin.Context, getFunc eventsFunc) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, limit, getFunc)
	} else {
		resp, err := getFunc(ctx.Params, limit, time.Time{})
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (h *EventsHandlers) AllNamespaceResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(true, true)...)
}

func (h *EventsHandlers) SelectedNamespaceResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(false, true, strings.Split(ctx.Query("res"), ",")...)...)
}

func (h *EventsHandlers) AllResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(true, false)...)
}

func (h *EventsHandlers) SelectedResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(false, false, strings.Split(ctx.Query("res"), ",")...)...)
}

func (h *EventsHandlers) GetEventsFuncs(all bool, ns bool, events ...string) (eventFuncs []eventsFunc) {
	var getMap = map[string]eventsFunc{}
	if ns {
		getMap = map[string]eventsFunc{
			"ns":         h.GetNamespaceChanges,
			"deploy":     h.GetNamespaceDeploymentsChanges,
			"svc":        h.GetNamespaceServicesChanges,
			"ingress":    h.GetNamespaceIngressesChanges,
			"cm":         h.GetNamespaceConfigMapsChanges,
			"secret":     h.GetNamespaceSecretsChanges,
			"pvc":        h.GetNamespacePVCsChanges,
			"events-pod": h.GetNamespacePodsEvents,
			"events-pvc": h.GetNamespacePVCsEvents,
		}
	} else {
		getMap = map[string]eventsFunc{
			"ns":         h.GetAllNamespacesChanges,
			"deploy":     h.GetAllNamespacesDeploymentsChanges,
			"svc":        h.GetAllNamespacesServicesChanges,
			"ingress":    h.GetAllNamespacesIngressesChanges,
			"cm":         h.GetAllNamespacesConfigMapsChanges,
			"secret":     h.GetAllNamespacesSecretsChanges,
			"pvc":        h.GetAllNamespacesPVCsChanges,
			"events-pod": h.GetAllNamespacesPodsEvents,
			"events-pvc": h.GetAllNamespacesPVCsEvents,
		}
	}

	if all {
		for _, newFunc := range getMap {
			eventFuncs = append(eventFuncs, newFunc)
		}
	} else {
		for _, event := range events {
			newFunc, ok := getMap[event]
			if ok {
				eventFuncs = append(eventFuncs, newFunc)
			}
		}
	}
	return
}
