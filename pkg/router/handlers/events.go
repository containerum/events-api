package handlers

import (
	"net/http"

	"github.com/containerum/events-api/pkg/eaerrors"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /events/namespaces/{namespace}/pods/{pod} Events GetPodEventsList
// Get pod events.
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
//  - name: pod
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetPodEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPodEvents)
}

// swagger:operation GET /events/namespaces/{namespace}/pods Events GetNamespacePodsEventsList
// Get namespace pods events.
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetNamespacePodsEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePodsEvents)
}

// swagger:operation GET /events/namespaces/{namespace}/pvc/{pvc} Events GetPVCEventsList
// Get PVC events.
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
//  - name: pvc
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetPVCEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPVCEvents)
}

// swagger:operation GET /events/namespaces/{namespace}/pvc Events GetNamespacePVCsEventsList
// Get namespace PVCs events.
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetNamespacePVCsEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePVCsEvents)
}

// swagger:operation GET /events/containerum/users Events GetUsersEventsList
// Get users events.
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetUsersEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetUsersEvents)
}

// swagger:operation GET /events/containerum/system Events GetSystemEventsList
// Get system events.
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
//  '200':
//    description: events list
//    schema:
//      $ref: '#/definitions/EventsList'
//  '101':
//    description: websocket response
//    schema:
//      $ref: '#/definitions/EventsList'
//  default:
//    $ref: '#/responses/error'
func (h *EventsHandlers) GetSystemEventsListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetSystemEvents)
}

// swagger:operation POST /events/containerum/users Events AddUserEvent
// Add user event.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: body
//    in: body
//    schema:
//      $ref: '#/definitions/Event'
// responses:
//  '202':
//    description: event added
//  default:
//    $ref: '#/responses/error'
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

// swagger:operation POST /events/containerum/system Events AddSystemEvent
// Add system event.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UserRoleHeader'
//  - $ref: '#/parameters/UserIDHeader'
//  - name: body
//    in: body
//    schema:
//      $ref: '#/definitions/Event'
// responses:
//  '202':
//    description: event added
//  default:
//    $ref: '#/responses/error'
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
