package eventfilter

import (
	"runtime"
	"sync"

	"github.com/containerum/kube-client/pkg/model"
)

func Or(predicates ...Predicate) Predicate {
	return func(event model.Event) bool {
		for _, predicate := range predicates {
			if predicate(event) {
				return true
			}
		}
		return false
	}
}

func And(predicates ...Predicate) Predicate {
	return func(event model.Event) bool {
		for _, predicate := range predicates {
			if !predicate(event) {
				return false
			}
		}
		return true
	}
}

func Not(predicate Predicate) Predicate {
	return func(event model.Event) bool {
		return !predicate(event)
	}
}

type Predicate func(model.Event) bool

func (predicate Predicate) Filter(events []model.Event) []model.Event {
	var filtered = make([]model.Event, 0, len(events))
	for _, event := range events {
		if predicate(event) {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

func (predicate Predicate) MatchAll(events []model.Event) bool {
	for _, event := range events {
		if !predicate(event) {
			return false
		}
	}
	return true
}

// Spawns n (if n>0, else runtime.NumCPU()) filtering workers
func (predicate Predicate) Pipe(n int, in <-chan model.Event) <-chan model.Event {
	if n <= 0 {
		n = runtime.NumCPU()
	}
	var out = make(chan model.Event, n)
	var wg = &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for event := range in {
				if predicate(event) {
					out <- event
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
