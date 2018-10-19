package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/containerum/events-api/pkg/eaerrors"
	"github.com/sirupsen/logrus"

	"github.com/containerum/events-api/pkg/model"
	"github.com/containerum/utils/httputil"

	m "github.com/containerum/events-api/pkg/router/middleware"
	"github.com/containerum/events-api/pkg/server"
	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	server.EventsActions
	*m.TranslateValidate
	DBPeriod time.Duration
}

func handleResourceChangesEvents(h *EventsHandlers, ctx *gin.Context, getFunc model.EventsFunc) {
	params := createParams(ctx)

	if _, ws := ctx.GetQuery("ws"); ws {
		withWS(ctx, params, h.DBPeriod, getFunc)
	} else {
		resp, err := getFunc(params)
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
	withWS(ctx, createParams(ctx), h.DBPeriod, h.getEventsFuncs(true, true)...)
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
	withWS(ctx, createParams(ctx), h.DBPeriod, h.getEventsFuncs(false, true, strings.Split(ctx.Query("res"), ",")...)...)
}

// swagger:operation GET /namespaces/ AllEvents AllNamespaceResourcesChangesEventsPaginatedHandler
// Get selected events in namespace.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: page
//    in: query
//    type: int
//    required: false
//  - name: page_size
//    in: query
//    type: integer
//    required: false
// responses:
//  '200':
//    description: page with events
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) AllNamespaceResourcesChangesEventsPaginatedHandler(ctx *gin.Context) {
	var params = createParams(ctx)
	var events, getPaginatedEventsErr = h.GetPaginatedEvents(params)
	if getPaginatedEventsErr != nil {
		logrus.WithError(getPaginatedEventsErr).Errorf("unable to get paginated event list: %v", params)
		ctx.AbortWithError(http.StatusInternalServerError, eaerrors.ErrInternal().AddDetails(
			"unable to get paginated event list",
			fmt.Sprint(params)))
	}
	ctx.JSON(http.StatusOK, events)
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
	withWS(ctx, createParams(ctx), h.DBPeriod, h.getEventsFuncs(true, false)...)
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
	withWS(ctx, createParams(ctx), h.DBPeriod, h.getEventsFuncs(false, false, strings.Split(ctx.Query("res"), ",")...)...)
}

func (h *EventsHandlers) getEventsFuncs(all, ns bool, events ...string) (eventFuncs []model.EventsFunc) {
	var getMap map[string]model.EventsFunc
	if ns {
		getMap = map[string]model.EventsFunc{
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
		getMap = map[string]model.EventsFunc{
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

func createParams(ctx *gin.Context) model.FuncParams {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	startTime, _ := time.Parse(time.RFC3339, ctx.Query("time"))
	isAdmin := httputil.MustGetUserRole(ctx.Request.Context()) == m.RoleAdmin
	var page, _ = strconv.Atoi(ctx.Query("page"))
	var pageSize, _ = strconv.Atoi(ctx.Query("page_size"))

	var namespaces []string
	if !isAdmin {
		nsList := ctx.MustGet(m.UserNamespaces).(*m.UserHeaderDataMap)
		for _, n := range *nsList {
			namespaces = append(namespaces, n.ID)
		}
	}

	return model.FuncParams{
		Params:         ctx.Params,
		Limit:          limit,
		StartTime:      startTime,
		UserAdmin:      isAdmin,
		UserNamespaces: namespaces,
		Page:           page,
		PageSize:       pageSize,
	}
}
