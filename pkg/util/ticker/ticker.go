package ticker

import (
	"sync"
	"time"
)

type Ticker struct {
	events chan time.Time
	stop   chan struct{}
	step   time.Duration
	once   *sync.Once
}

func NewTicker(step time.Duration) Ticker {
	return Ticker{
		events: make(chan time.Time),
		stop:   make(chan struct{}),
		once:   &sync.Once{},
		step:   step,
	}
}

func (ticker Ticker) Start() {
	ticker.tick()
}

func (ticker Ticker) Stop() {
	ticker.once.Do(func() {
		close(ticker.stop)
	})
}

func (ticker Ticker) Closed() <-chan struct{} {
	return ticker.stop
}

func (ticker Ticker) Ticks() <-chan time.Time {
	return ticker.events
}

func (ticker Ticker) tick() {
	select {
	case ticker.events <- time.Now():
		time.AfterFunc(ticker.step, ticker.tick)
	case <-ticker.stop:
		close(ticker.events)
		return
	}
}
