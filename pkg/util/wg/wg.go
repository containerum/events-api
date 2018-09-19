package wg

import (
	"sync"
	"sync/atomic"
)

type WG struct {
	wg      *sync.WaitGroup
	done    int64
	allDone chan struct{}
}

func NewWG(n int) *WG {
	var wg = &sync.WaitGroup{}
	wg.Add(n)
	var allDone = make(chan struct{})
	go func() {
		wg.Done()
		close(allDone)
	}()
	return &WG{
		wg:      wg,
		allDone: allDone,
		done:    int64(n),
	}
}

func (doner *WG) Done() {
	atomic.AddInt64(&doner.done, -1)
	doner.wg.Done()
}

func (doner *WG) Left() int64 {
	return atomic.LoadInt64(&doner.done)
}

func (doner *WG) Wait() <-chan struct{} {
	return doner.allDone
}
