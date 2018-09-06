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

func NewEventsActionsImpl(mongo *db.MongoStorage) *EventsActionsImpl {
	return &EventsActionsImpl{
		mongo: mongo,
		log:   cherrylog.NewLogrusAdapter(logrus.WithField("component", "domain_actions")),
	}
}

func (ea *EventsActionsImpl) GetPodEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("pod"), db.PodEventsCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.PodEventsCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPVCEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("pvc"), db.PVCEventsCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.PVCEventsCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}
