package impl

import (
	"time"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

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
	events, err := ea.mongo.GetEventsList(params.ByName("namespace"), params.ByName("pod"), model.TypePod, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsInNamespaceList(params.ByName("namespace"), model.TypePod, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPVCEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsList(params.ByName("namespace"), params.ByName("pvc"), model.TypeVolume, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsEvents(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsInNamespaceList(params.ByName("namespace"), model.TypeVolume, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) AddUserEvent(event model.Event) error {
	event.DateAdded = time.Now().Format(time.RFC3339)
	event.ResourceType = model.TypeUser
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.UserCollection, event)
}

func (ea *EventsActionsImpl) AddSystemEvent(event model.Event) error {
	event.DateAdded = time.Now().Format(time.RFC3339)
	event.ResourceType = model.TypeSystem
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.SystemCollection, event)
}
