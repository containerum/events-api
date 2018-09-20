package impl

import (
	"time"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

func (ea *EventsActionsImpl) GetNamespaceChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting namespace changes")
	changes, err := ea.mongo.GetNamespaceChangesList(ns, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetDeploymentChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	deploy := params.ByName("deployment")
	ea.log.WithField("namespace", ns).WithField("deployment", deploy).Debugln("Getting deployment changes")
	changes, err := ea.mongo.GetChangesList(ns, deploy, mongodb.DeploymentCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting deployments changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.DeploymentCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetServiceChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	svc := params.ByName("service")
	ea.log.WithField("namespace", ns).WithField("service", svc).Debugln("Getting service changes")
	changes, err := ea.mongo.GetChangesList(ns, svc, mongodb.ServiceCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceServicesChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting services changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.ServiceCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetIngressChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ingr := params.ByName("ingress")
	ea.log.WithField("namespace", ns).WithField("ingress", ingr).Debugln("Getting ingress changes")
	changes, err := ea.mongo.GetChangesList(ns, ingr, mongodb.IngressCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceIngressesChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting ingresses changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.IngressCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetPVCChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	pvc := params.ByName("pvc")
	ea.log.WithField("namespace", ns).WithField("pvc", pvc).Debugln("Getting PVC changes")
	changes, err := ea.mongo.GetChangesList(ns, pvc, mongodb.PVCCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting PVCs changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.PVCCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetSecretChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	secret := params.ByName("pvc")
	ea.log.WithField("namespace", ns).WithField("secret", secret).Debugln("Getting secret changes")
	changes, err := ea.mongo.GetChangesList(ns, secret, mongodb.SecretsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceSecretsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting secrets changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.SecretsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetConfigMapChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	cm := params.ByName("configmap")
	ea.log.WithField("namespace", ns).WithField("configmap", cm).Debugln("Getting configmap changes")
	changes, err := ea.mongo.GetChangesList(ns, cm, mongodb.ConfigMapsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceConfigMapsChanges(params gin.Params, limit int, startTime time.Time) (*model.EventsList, error) {
	ns := params.ByName("namespace")
	ea.log.WithField("namespace", ns).Debugln("Getting configmaps changes")
	changes, err := ea.mongo.GetChangesInNamespaceList(ns, mongodb.ConfigMapsCollection, limit, startTime)
	if err != nil {
		return nil, err
	}
	return &model.EventsList{Events: changes}, nil
}
