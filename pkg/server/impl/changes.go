package impl

import (
	"time"

	"github.com/containerum/events-api/pkg/db"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

func (ea *EventsActionsImpl) GetNamespaceChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespaceChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetDeploymentChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("deployment"), db.DeploymentCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.DeploymentCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetServiceChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("service"), db.ServiceCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceServicesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.ServiceCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetIngressChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("ingress"), db.IngressCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceIngressesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.IngressCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetPVCChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("pvc"), db.PVCCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), db.PVCCollection, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}
