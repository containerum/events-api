package server

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

type EventsActions interface {
	GetDeploymentEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespaceDeploymentsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)

	GetPodEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
	GetNamespacePodsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error)
}
