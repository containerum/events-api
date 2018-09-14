package impl

import (
	"time"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

func (ea *EventsActionsImpl) GetNamespaceChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetNamespaceChangesList(params.ByName("namespace"), startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetDeploymentChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("deployment"), mongodb.DeploymentCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.DeploymentCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetServiceChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("service"), mongodb.ServiceCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceServicesChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.ServiceCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetIngressChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("ingress"), mongodb.IngressCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceIngressesChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.IngressCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetPVCChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("pvc"), mongodb.PVCCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.PVCCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetSecretChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("secret"), mongodb.SecretsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceSecretsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.SecretsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetConfigMapChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesList(params.ByName("namespace"), params.ByName("configmap"), mongodb.ConfigMapsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceConfigMapsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	changes, err := ea.mongo.GetChangesInNamespaceList(params.ByName("namespace"), mongodb.ConfigMapsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}
