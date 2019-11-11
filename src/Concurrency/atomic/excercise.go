package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// create race condition.All the go routines will try to access the variable at a same time.

func main() {
	var incrementor int64
	var gr sync.WaitGroup
	gs := 100
	gr.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//Adding the delta to incrementor
			atomic.AddInt64(&incrementor, 1)

			runtime.Gosched()

			gr.Done()
		}()

	}
	gr.Wait()
	println(incrementor)
}
