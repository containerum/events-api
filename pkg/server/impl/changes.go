package impl

import (
	"github.com/containerum/events-api/pkg/model"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

	kubeModel "github.com/containerum/kube-client/pkg/model"
)

func (ea *EventsActionsImpl) GetNamespaceChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting namespace changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetNamespacesChangesList(params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting all namespaces changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllNamespacesChangesList(params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetNamespacesChangesList(params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetDeploymentChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	deploy := params.Params.ByName("deployment")

	ea.log.WithField("namespace", ns).WithField("deployment", deploy).Debugln("Getting deployment changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, deploy, mongodb.DeploymentCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceDeploymentsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting deployments changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.DeploymentCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesDeploymentsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting deployments changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.DeploymentCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.DeploymentCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetServiceChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	svc := params.Params.ByName("service")

	ea.log.WithField("namespace", ns).WithField("service", svc).Debugln("Getting service changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, svc, mongodb.ServiceCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceServicesChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting services changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.ServiceCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesServicesChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting services changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.ServiceCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.ServiceCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetIngressChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	ingr := params.Params.ByName("ingress")

	ea.log.WithField("namespace", ns).WithField("ingress", ingr).Debugln("Getting ingress changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, ingr, mongodb.IngressCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceIngressesChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting ingresses changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.IngressCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesIngressesChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting ingresses changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.IngressCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.IngressCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetPVCChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	pvc := params.Params.ByName("pvc")

	ea.log.WithField("namespace", ns).WithField("pvc", pvc).Debugln("Getting PVC changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, pvc, mongodb.PVCCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespacePVCsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting PVCs changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.PVCCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesPVCsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting PVCs changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.PVCCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.PVCCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetSecretChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	secret := params.Params.ByName("pvc")

	ea.log.WithField("namespace", ns).WithField("secret", secret).Debugln("Getting secret changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, secret, mongodb.SecretsCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceSecretsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting secrets changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.SecretsCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesSecretsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting secrets changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.SecretsCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.SecretsCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetConfigMapChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")
	cm := params.Params.ByName("configmap")

	ea.log.WithField("namespace", ns).WithField("configmap", cm).Debugln("Getting configmap changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesList(ns, cm, mongodb.ConfigMapsCollection, params.Limit, params.StartTime)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetNamespaceConfigMapsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ns := params.Params.ByName("namespace")

	ea.log.WithField("namespace", ns).Debugln("Getting configmaps changes")
	ns = checkNamespacePermissions(params.UserAdmin, ns, params.UserNamespaces)

	changes, err := ea.mongo.GetChangesInNamespacesList(mongodb.ConfigMapsCollection, params.Limit, params.StartTime, ns)
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}

func (ea *EventsActionsImpl) GetAllNamespacesConfigMapsChanges(params model.FuncParams) (*kubeModel.EventsList, error) {
	ea.log.Debugln("Getting configmaps changes")

	var changes []kubeModel.Event
	var err error
	if params.UserAdmin {
		changes, err = ea.mongo.GetAllChangesList(mongodb.ConfigMapsCollection, params.Limit, params.StartTime)
	} else {
		changes, err = ea.mongo.GetChangesInNamespacesList(mongodb.ConfigMapsCollection, params.Limit, params.StartTime, params.UserNamespaces...)
	}
	if err != nil {
		return nil, err
	}
	return &kubeModel.EventsList{Events: changes}, nil
}
