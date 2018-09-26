package handlers

import (
	"github.com/gin-gonic/gin"
)

// swagger:operation GET /changes/namespaces/{namespace} Changes GetNamespaceChangesList
// Get namespace changes.
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
func (h *EventsHandlers) GetNamespaceChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/deployments/{deployment} Changes GetDeploymentChangesList
// Get deployment changes.
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
//  - name: deployment
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
func (h *EventsHandlers) GetDeploymentChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetDeploymentChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/deployments Changes GetNamespaceDeploymentsChangesList
// Get namespace deployments changes.
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
func (h *EventsHandlers) GetNamespaceDeploymentsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceDeploymentsChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/services/{service} Changes GetServiceChangesList
// Get service changes.
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
//  - name: service
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
func (h *EventsHandlers) GetServiceChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetServiceChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/services Changes GetNamespaceServicesChangesList
// Get namespace services changes.
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
func (h *EventsHandlers) GetNamespaceServicesChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceServicesChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/ingresses/{ingress} Changes GetIngressChangesList
// Get ingress changes.
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
//  - name: ingress
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
func (h *EventsHandlers) GetIngressChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetIngressChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/ingresses Changes GetNamespaceIngressesChangesList
// Get namespace ingresses changes.
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
func (h *EventsHandlers) GetNamespaceIngressesChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceIngressesChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/pvc/{pvc} Changes GetPVCChangesList
// Get PVC changes.
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
func (h *EventsHandlers) GetPVCChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPVCChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/pvc Changes GetNamespacePVCsChangesList
// Get namespace PVCs changes.
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
func (h *EventsHandlers) GetNamespacePVCsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePVCsChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/secrets/{secret} Changes GetSecretChangesList
// Get secret changes.
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
//  - name: secret
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
func (h *EventsHandlers) GetSecretChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetSecretChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/secrets Changes GetNamespaceSecretsChangesList
// Get namespace secrets changes.
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
func (h *EventsHandlers) GetNamespaceSecretsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceSecretsChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/configmaps/{configmap} Changes GetConfigMapChangesList
// Get configmap changes.
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
//  - name: configmap
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
func (h *EventsHandlers) GetConfigMapChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetConfigMapChanges)
}

// swagger:operation GET /changes/namespaces/{namespace}/configmaps Changes GetNamespaceConfigMapsChangesList
// Get namespace configmaps changes.
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
func (h *EventsHandlers) GetNamespaceConfigMapsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceConfigMapsChanges)
}
