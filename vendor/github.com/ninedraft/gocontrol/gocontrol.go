package gocontrol

import (
	"sync"
	"sync/atomic"
)

type Guard struct {
	counter int64
	wg      sync.WaitGroup
}

func (guard *Guard) Go() func() {
	atomic.AddInt64(&guard.counter, 1)
	guard.wg.Add(1)
	return func() {
		guard.wg.Done()
		atomic.AddInt64(&guard.counter, -1)
	}
}

func (guard *Guard) AliveN() int64 {
	return atomic.LoadInt64(&guard.counter)
}

func (guard *Guard) Wait() {
	guard.wg.Wait()
}

// Await all goroutines and call closure
//  defer guard.Await(func() {
//	 fmt.Println("DONE")
//  })
func (guard *Guard) Await(op func()) {
	guard.Wait()
	op()
}
