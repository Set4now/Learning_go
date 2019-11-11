package main

import (
	"runtime"
	"sync"
)

func main() {
	incrementor := 0
	var gr sync.WaitGroup
	gs := 100
	gr.Add(gs)
	//Create a variable of value of type Mutex from sync package
	var lock sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// Locking the data
			lock.Lock()
			a := incrementor

			runtime.Gosched()
			a++

			incrementor = a
			// Unlocking the data
			lock.Unlock()
			gr.Done()
		}()

	}
	gr.Wait()
	println(incrementor)
}
