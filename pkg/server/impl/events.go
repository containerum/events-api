package impl

import (
	"context"

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

func (ea *EventsActionsImpl) GetDeploymentEvents(ctx context.Context, namespace, deployment string) (*model.EventsList, error) {
	events, err := ea.mongo.GetDeploymentEventsList(namespace, deployment)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsEvents(ctx context.Context, namespace string) (*model.EventsList, error) {
	events, err := ea.mongo.GetNamespaceDeploymentsEventsList(namespace)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPodEvents(ctx context.Context, namespace, pod string) (*model.EventsList, error) {
	events, err := ea.mongo.GetPodEventsList(namespace, pod)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(ctx context.Context, namespace string) (*model.EventsList, error) {
	events, err := ea.mongo.GetNamespacePodsEventsList(namespace)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}
