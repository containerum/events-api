package server

import (
	"github.com/containerum/events-api/pkg/model"
	kubeModel "github.com/containerum/kube-client/pkg/model"
)

type EventsActions interface {
	GetNamespaceChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetDeploymentChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespaceDeploymentsChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesDeploymentsChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetServiceChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespaceServicesChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesServicesChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetIngressChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespaceIngressesChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesIngressesChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetPVCChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespacePVCsChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesPVCsChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetSecretChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespaceSecretsChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesSecretsChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetConfigMapChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespaceConfigMapsChanges(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesConfigMapsChanges(params model.FuncParams) (*kubeModel.EventsList, error)

	GetPodEvents(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespacePodsEvents(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesPodsEvents(params model.FuncParams) (*kubeModel.EventsList, error)

	GetPVCEvents(params model.FuncParams) (*kubeModel.EventsList, error)
	GetNamespacePVCsEvents(params model.FuncParams) (*kubeModel.EventsList, error)
	GetAllNamespacesPVCsEvents(params model.FuncParams) (*kubeModel.EventsList, error)

	GetAllNodesEvents(params model.FuncParams) (*kubeModel.EventsList, error)

	GetUsersEvents(params model.FuncParams) (*kubeModel.EventsList, error)
	GetSystemEvents(params model.FuncParams) (*kubeModel.EventsList, error)

	AddUserEvent(event kubeModel.Event) error
	AddSystemEvent(event kubeModel.Event) error
}
