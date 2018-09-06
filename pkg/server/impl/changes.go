package impl

import (
	"time"

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
	changes, err := ea.mongo.GetDeploymentChangesList(params.ByName("namespace"), params.ByName("deployment"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespaceDeploymentsChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetServiceChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetServiceChangesList(params.ByName("namespace"), params.ByName("service"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceServicesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespaceServicesChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetIngressChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetIngressChangesList(params.ByName("namespace"), params.ByName("ingress"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceIngressesChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespaceIngressesChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetPVCChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetPVCChangesList(params.ByName("namespace"), params.ByName("pvc"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsChanges(params gin.Params, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespacePVCsChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}
