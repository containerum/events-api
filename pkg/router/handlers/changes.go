package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *EventsHandlers) GetNamespaceChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceChanges)
}

func (h *EventsHandlers) GetDeploymentChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetDeploymentChanges)
}

func (h *EventsHandlers) GetNamespaceDeploymentsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceDeploymentsChanges)
}

func (h *EventsHandlers) GetServiceChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetServiceChanges)
}

func (h *EventsHandlers) GetNamespaceServicesChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceServicesChanges)
}

func (h *EventsHandlers) GetIngressChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetIngressChanges)
}

func (h *EventsHandlers) GetNamespaceIngressesChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceIngressesChanges)
}

func (h *EventsHandlers) GetPVCChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetPVCChanges)
}

func (h *EventsHandlers) GetNamespacePVCsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespacePVCsChanges)
}

func (h *EventsHandlers) GetSecretChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetSecretChanges)
}

func (h *EventsHandlers) GetNamespaceSecretsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceSecretsChanges)
}

func (h *EventsHandlers) GetConfigMapChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetConfigMapChanges)
}

func (h *EventsHandlers) GetNamespaceConfigMapsChangesListHandler(ctx *gin.Context) {
	handleResourceChangesEvents(h, ctx, h.GetNamespaceConfigMapsChanges)
}
