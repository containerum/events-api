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
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
}

type eventsFunc func(params gin.Params, limit int, startFrom time.Time) (*model.EventsList, error)

func handleResourceChangesEvents(h *EventsHandlers, ctx *gin.Context, getFunc eventsFunc) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, limit, startTime, getFunc)
	} else {
		resp, err := getFunc(ctx.Params, limit, startTime)
		if err != nil {
			ctx.AbortWithStatusJSON(h.HandleError(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

// swagger:operation GET /namespaces/{namespace}/all AllEvents AllNamespaceResourcesChangesEvents
// Get all events in namespace.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: namespace
//    in: path
//    type: string
//    required: true
//  - name: ws
//    in: query
//    type: string
//    required: false
//  - name: limit
//    in: query
//    type: string
//    required: false
//  - name: time
//    in: query
//    type: string
//    required: false
// responses:
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) AllNamespaceResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	withWS(ctx, limit, startTime, h.getEventsFuncs(true, true)...)
}

// swagger:operation GET /namespaces/{namespace}/selected AllEvents SelectedNamespaceResourcesChangesEvents
// Get selected events in namespace.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: namespace
//    in: path
//    type: string
//    required: true
//  - name: ws
//    in: query
//    type: string
//    required: false
//  - name: limit
//    in: query
//    type: string
//    required: false
//  - name: res
//    in: query
//    type: string
//    required: false
//  - name: time
//    in: query
//    type: string
//    required: false
// responses:
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) SelectedNamespaceResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	withWS(ctx, limit, startTime, h.getEventsFuncs(false, true, strings.Split(ctx.Query("res"), ",")...)...)
}

// swagger:operation GET /all AllEvents AllResourcesChangesEvents
// Get all events.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: ws
//    in: query
//    type: string
//    required: false
//  - name: limit
//    in: query
//    type: string
//    required: false
//  - name: time
//    in: query
//    type: string
//    required: false
// responses:
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) AllResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	withWS(ctx, limit, startTime, h.getEventsFuncs(true, false)...)
}

// swagger:operation GET /selected AllEvents SelectedResourcesChangesEvents
// Get selected events.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: ws
//    in: query
//    type: string
//    required: false
//  - name: limit
//    in: query
//    type: string
//    required: false
//  - name: res
//    in: query
//    type: string
//    required: false
//  - name: time
//    in: query
//    type: string
//    required: false
// responses:
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) SelectedResourcesChangesEventsHandler(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	withWS(ctx, limit, startTime, h.getEventsFuncs(false, false, strings.Split(ctx.Query("res"), ",")...)...)
}

func (h *EventsHandlers) getEventsFuncs(all, ns bool, events ...string) (eventFuncs []eventsFunc) {
	var getMap map[string]eventsFunc
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
	return eventFuncs
}
