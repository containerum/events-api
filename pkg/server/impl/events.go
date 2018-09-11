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

func (ea *EventsActionsImpl) GetPodEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsList(params.ByName("namespace"), params.ByName("pod"), model.TypePod, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsInNamespaceList(params.ByName("namespace"), model.TypePod, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPVCEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsList(params.ByName("namespace"), params.ByName("pvc"), model.TypeVolume, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetEventsInNamespaceList(params.ByName("namespace"), model.TypeVolume, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) AddUserEvent(event model.Event) error {
	event.DateAdded = time.Now()
	event.ResourceType = model.TypeUser
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.UserCollection, event)
}

func (ea *EventsActionsImpl) AddSystemEvent(event model.Event) error {
	event.DateAdded = time.Now()
	event.ResourceType = model.TypeSystem
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.SystemCollection, event)
}

func (ea *EventsActionsImpl) GetUsersEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetUsersEventsList(limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetSystemEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	events, err := ea.mongo.GetSystemEventsList(limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}
