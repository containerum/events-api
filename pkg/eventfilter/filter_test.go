package eventfilter

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"

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

func BenchmarkFilter(b *testing.B) {
	var N = uint64(1000000)
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	var events = generateEvents(rnd, N)
	var filter = Or(
		And(EqType(model.TypePod), EqKind(model.EventInfo)),
		func(event model.Event) bool {
			time.Sleep(time.Nanosecond)
			return false
		},
		EqKind(model.EventError),
		EqKind(model.EventWarning),
	)
	b.Run("plain filter", func(b *testing.B) {
		var filteredEvents = make([]model.Event, 0, N)
		b.StartTimer()
		for _, event := range filter.Filter(events) {
			filteredEvents = append(filteredEvents, event)
		}
		b.StopTimer()
		b.ReportAllocs()
		if testing.Verbose() {
			b.Logf("%v events", filteredEvents)
		}
	})
	for i := 0; i < 1; i++ {
		for _, workersN := range []int{1, 2, 3, runtime.NumCPU()} {
			b.Run(fmt.Sprintf("filter %0d workers", workersN), func(b *testing.B) {
				var eventChan = filter.Pipe(workersN, eventSource(rnd, N))
				var filteredEvents = make([]model.Event, 0, N)
				b.StartTimer()
				for event := range eventChan {
					filteredEvents = append(filteredEvents, event)
				}
				b.StopTimer()
				b.ReportAllocs()
				if testing.Verbose() {
					b.Logf("%v events", filteredEvents)
				}
			})
		}
	}
}

func eventSource(rnd *rand.Rand, n uint64) <-chan model.Event {
	var events = make(chan model.Event, n)
	go func() {
		for i := uint64(0); i < n; i++ {
			events <- createEvent(rnd)
		}
		close(events)
	}()
	return events
}

var (
	kinds = []model.EventKind{
		model.EventInfo,
		model.EventWarning,
		model.EventError,
	}
	kindsN    = len(kinds)
	resources = []model.ResourceType{
		model.TypeDeployment,
		model.TypeIngress,
		model.TypeStorage,
		model.TypePod,
		model.ImportSuccessfulMessage,
		model.TypeConfigMap,
		model.TypeNamespace,
		model.TypeNode,
		model.TypeSecret,
		model.TypeVolume,
		model.TypeUser,
		model.TypeSystem,
	}
	resourcesN = len(resources)
	messages   = []string{
		"The world is indeed comic, but the joke is on mankind.",
		"Pleasure to me is wonder—the unexplored, the unexpected, the thing that is hidden and the changeless thing that lurks behind superficial mutability.",
		"The oldest and strongest emotion of mankind is fear, and the oldest and strongest kind of fear is fear of the unknown",
		"Almost nobody dances sober, unless they happen to be insane.",
		"From even the greatest of horrors irony is seldom absent.",
		"That is not dead which can eternal lie,\nAnd with strange aeons even death may die.",
		"The most merciful thing in the world, I think, is the inability of the human mind to correlate all its contents. We live on a placid island of ignorance in the midst of black seas of the infinity, and it was not meant that we should voyage far.",
	}
	messagesN = len(messages)
)

func createEvent(rnd *rand.Rand) model.Event {
	return model.Event{
		Kind:         kinds[rnd.Intn(kindsN)],
		ResourceType: resources[rnd.Intn(resourcesN)],
		Message:      messages[rnd.Intn(messagesN)],
	}
}

func generateEvents(rnd *rand.Rand, n uint64) []model.Event {
	var events = make([]model.Event, 0, n)
	for i := uint64(0); i < n; i++ {
		events = append(events, createEvent(rnd))
	}
	return events
}
