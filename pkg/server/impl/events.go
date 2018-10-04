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
	ns := params.ByName("namespace")
	pod := params.ByName("pod")
	ea.log.WithField("namespace", ns).WithField("pod", pod).Debugln("Getting pod events")
	events, err := ea.mongo.GetEventsList(ns, pod, model.TypePod, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting pods events")
	events, err := ea.mongo.GetEventsInNamespaceList(ns, model.TypePod, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesPodsEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ea.log.Debugln("Getting pods events")
	events, err := ea.mongo.GetAllEventsList(model.TypePod, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPVCEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	pvc := params.ByName("pvc")
	ea.log.WithField("namespace", ns).Debugln("Getting PVC events")
	events, err := ea.mongo.GetEventsList(ns, pvc, model.TypeVolume, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting PVCs events")
	events, err := ea.mongo.GetEventsInNamespaceList(ns, model.TypeVolume, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesPVCsEvents(_ gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ea.log.Debugln("Getting PVCs events")
	events, err := ea.mongo.GetAllEventsList(model.TypeVolume, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNodesEvents(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ea.log.Debugln("Getting nodes events")
	events, err := ea.mongo.GetAllEventsList(model.TypeNode, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetUsersEvents(_ gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ea.log.Debugln("Getting users events")
	events, err := ea.mongo.GetUsersEventsList(limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetSystemEvents(_ gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ea.log.Debugln("Getting system events")
	events, err := ea.mongo.GetSystemEventsList(limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) AddUserEvent(event model.Event) error {
	ea.log.Debugln("Adding user event")
	event.DateAdded = time.Now()
	event.ResourceType = model.TypeUser
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.UserCollection, event)
}

func (ea *EventsActionsImpl) AddSystemEvent(event model.Event) error {
	ea.log.Debugln("Adding system event")
	event.DateAdded = time.Now()
	event.ResourceType = model.TypeSystem
	if event.Kind == "" {
		event.Kind = model.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.SystemCollection, event)
}
