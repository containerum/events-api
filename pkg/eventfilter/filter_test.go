package eventfilter

import (
	"testing"

	"github.com/containerum/kube-client/pkg/model"
)

func TestFilter(test *testing.T) {
	var events = []model.Event{
		{Kind: model.EventInfo, ResourceType: model.TypePod, Message: "pod info message"},
		{Kind: model.EventInfo, ResourceType: model.TypeStorage, Message: "storage info message"},
		{Kind: model.EventInfo, ResourceType: model.TypePod, Message: "pod info message"},
		{Kind: model.EventWarning, ResourceType: model.TypeIngress, Message: "invalid cert format"},
		{Kind: model.EventWarning, ResourceType: model.TypeDeployment, Message: "unable to get Moon Rabbit"},
	}
	var filtered = Or(
		EqKind(model.EventError),
		EqKind(model.EventWarning),
		And(
			EqKind(model.EventInfo),
			EqType(model.TypeStorage)),
	).Filter(events)
	if len(filtered) != 3 {
		test.Fatal(filtered)
	}
	test.Log(filtered)
}
