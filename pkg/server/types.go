package server

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

type EventsActions interface {
	GetNamespaceChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetDeploymentChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespaceDeploymentsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetServiceChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespaceServicesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetIngressChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespaceIngressesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetPVCChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespacePVCsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetPodEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespacePodsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetPVCEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespacePVCsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetUsersEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetSystemEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)

	AddUserEvent(event model.Event) error
	AddSystemEvent(event model.Event) error
}
