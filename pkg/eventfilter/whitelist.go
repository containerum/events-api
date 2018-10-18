package eventfilter

import "github.com/containerum/kube-client/pkg/model"

type Whitelist struct {
	UserEvents  []string
	AdminEvents []string
}

func (whitelist Whitelist) UserEventsFilter() Predicate {
	return whitelist.filterByEvent(whitelist.UserEvents)
}

func (whitelist Whitelist) AdminEventsFilter() Predicate {
	return whitelist.filterByEvent(whitelist.AdminEvents)
}

func (whitelist Whitelist) filterByEvent(events []string) Predicate {
	var eventSet = make(map[string]struct{}, len(events))
	return func(event model.Event) bool {
		var _, ok = eventSet[event.Name]
		return ok
	}
}
