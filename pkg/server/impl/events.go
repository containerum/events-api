package impl

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/containerum/kube-client/pkg/model"

	"github.com/containerum/cherry/adaptors/cherrylog"
	"github.com/containerum/events-api/pkg/db"
	"github.com/sirupsen/logrus"
)

type EventsActionsImpl struct {
	mongo *db.MongoStorage
	log   *cherrylog.LogrusAdapter
}

func NewDomainActionsImpl(mongo *db.MongoStorage) *EventsActionsImpl {
	return &EventsActionsImpl{
		mongo: mongo,
		log:   cherrylog.NewLogrusAdapter(logrus.WithField("component", "domain_actions")),
	}
}

func (ea *EventsActionsImpl) GetDeploymentEvents(params gin.Params, starttime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetDeploymentEventsList(params.ByName("namespace"), params.ByName("deployment"), starttime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsEvents(params gin.Params, starttime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetNamespaceDeploymentsEventsList(params.ByName("namespace"), starttime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPodEvents(params gin.Params, starttime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetPodEventsList(params.ByName("namespace"), params.ByName("pod"), starttime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params gin.Params, starttime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetNamespacePodsEventsList(params.ByName("namespace"), starttime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}
