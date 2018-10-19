package eventfilter

import "github.com/containerum/kube-client/pkg/model"

func True(model.Event) bool {
	return true
}

func False(model.Event) bool {
	return false
}

func EqKind(kind model.EventKind) Predicate {
	return func(event model.Event) bool { return event.Kind == kind }
}

func MatchAnyKind(kinds ...model.EventKind) Predicate {
	return func(event model.Event) bool {
		for _, kind := range kinds {
			if kind == event.Kind {
				return true
			}
		}
		return false
	}
}

func EqType(resourceType model.ResourceType) Predicate {
	return func(event model.Event) bool {
		return event.ResourceType == resourceType
	}
}

func MatchAnyType(resourceTypes ...model.ResourceType) Predicate {
	return func(event model.Event) bool {
		for _, resourceType := range resourceTypes {
			if resourceType == event.ResourceType {
				return true
			}
		}
		return false
	}
}

func EqResourceName(name string) Predicate {
	return func(event model.Event) bool {
		return event.ResourceName == name
	}
}

func MatchResourceName(names ...string) Predicate {
	return func(event model.Event) bool {
		for _, name := range names {
			if name == event.ResourceName {
				return true
			}
		}
		return false
	}
}
