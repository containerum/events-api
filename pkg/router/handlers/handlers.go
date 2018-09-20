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

func (h *EventsHandlers) AllResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(true)...)
}

func (h *EventsHandlers) SelectedResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		logrus.Warn(err)
	}
	withWS(ctx, limit, h.GetEventsFuncs(false, strings.Split(ctx.Query("res"), ",")...)...)
}

func (h *EventsHandlers) GetEventsFuncs(all bool, events ...string) (eventFuncs []eventsFunc) {
	var getMap = map[string]eventsFunc{
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
