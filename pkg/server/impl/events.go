package impl

import (
	"time"

	"github.com/containerum/events-api/pkg/model"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

	kubeModel "github.com/containerum/kube-client/pkg/model"

	"github.com/containerum/cherry/adaptors/cherrylog"
	"github.com/containerum/events-api/pkg/db"
	log "github.com/sirupsen/logrus"
)

type EventsActionsImpl struct {
	mongo *db.MongoStorage
	log   *cherrylog.LogrusAdapter
}

func NewEventsActionsImpl(mongo *db.MongoStorage) *EventsActionsImpl {
	return &EventsActionsImpl{
		mongo: mongo,
		log:   cherrylog.NewLogrusAdapter(log.WithField("component", "domain_actions")),
	}
}

func (ea *EventsActionsImpl) GetPodEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	pod := params.Params.ByName("pod")

	ea.log.WithField("namespace", ns).WithField("pod", pod).Debugln("Getting pod events")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	events, err := ea.mongo.GetEventsList(ns, pod, kubeModel.TypePod, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePodsEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting pods events")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	events, err := ea.mongo.GetEventsInNamespacesList(kubeModel.TypePod, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesPodsEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting pods events")

	var events []kubeModel.Event
	var err error
	if params.UserAdmin {
		events, err = ea.mongo.GetAllEventsList(kubeModel.TypePod, params.Limit, params.StartTime)
	} else {
		events, err = ea.mongo.GetEventsInNamespacesList(kubeModel.TypePod, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetPVCEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	pvc := params.Params.ByName("pvc")

	ea.log.WithField("namespace", ns).Debugln("Getting PVC events")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	events, err := ea.mongo.GetEventsList(ns, pvc, kubeModel.TypeVolume, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting PVCs events")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	events, err := ea.mongo.GetEventsInNamespacesList(kubeModel.TypeVolume, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesPVCsEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting PVCs events")

	var events []kubeModel.Event
	var err error
	if params.UserAdmin {
		events, err = ea.mongo.GetAllEventsList(kubeModel.TypeVolume, params.Limit, params.StartTime)
	} else {
		events, err = ea.mongo.GetEventsInNamespacesList(kubeModel.TypeVolume, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetAllNodesEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting nodes events")
	events, err := ea.mongo.GetAllEventsList(kubeModel.TypeNode, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetUsersEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting users events")
	events, err := ea.mongo.GetUsersEventsList(params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) GetSystemEvents(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting system events")
	events, err := ea.mongo.GetSystemEventsList(params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: events}, nil
}

func (ea *EventsActionsImpl) AddUserEvent(event kubeModel.Event) error {
	ea.log.Debugln("Adding user event")
	event.DateAdded = time.Now()
	event.ResourceType = kubeModel.TypeUser
	if event.Kind == "" {
		event.Kind = kubeModel.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.UserCollection, event)
}

func (ea *EventsActionsImpl) AddSystemEvent(event kubeModel.Event) error {
	ea.log.Debugln("Adding system event")
	event.DateAdded = time.Now()
	event.ResourceType = kubeModel.TypeSystem
	if event.Kind == "" {
		event.Kind = kubeModel.EventInfo
	}
	return ea.mongo.AddContainerumEvent(mongodb.SystemCollection, event)
}
