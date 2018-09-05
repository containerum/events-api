package server

import (
	"context"

	"github.com/containerum/kube-client/pkg/model"
)

type EventsActions interface {
	GetDeploymentEvents(ctx context.Context, namespace, deployment string) (*model.EventsList, error)
	GetNamespaceDeploymentsEvents(ctx context.Context, namespace string) (*model.EventsList, error)

	GetPodEvents(ctx context.Context, namespace, pod string) (*model.EventsList, error)
	GetNamespacePodsEvents(ctx context.Context, namespace string) (*model.EventsList, error)
}
