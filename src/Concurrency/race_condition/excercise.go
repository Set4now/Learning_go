package main

import (
	"runtime"
	"sync"
)

// create race condition.All the go routines will try to access the variable at a same time.

func main() {
	incrementor := 0
	var gr sync.WaitGroup
	gs := 100
	gr.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			a := incrementor

			runtime.Gosched()
			a++

			incrementor = a

			gr.Done()
		}()

	}
	gr.Wait()
	println(incrementor)
}

//m-c02tm0amg8wn:Concurrency s0d011k$ go run --race race_condition/excercise.go
// ==================
// WARNING: DATA RACE
// Read at 0x00c000086000 by goroutine 6:
//   main.main.func1()
//       /Users/s0d011k/learning_go/src/github.com/Set4now/Concurrency/race_condition/excercise.go:16 +0x3c

// Previous write at 0x00c000086000 by goroutine 5:
//   main.main.func1()
//       /Users/s0d011k/learning_go/src/github.com/Set4now/Concurrency/race_condition/excercise.go:19 +0x5c

// Goroutine 6 (running) created at:
//   main.main()
//       /Users/s0d011k/learning_go/src/github.com/Set4now/Concurrency/race_condition/excercise.go:15 +0xe4

// Goroutine 5 (finished) created at:
//   main.main()
//       /Users/s0d011k/learning_go/src/github.com/Set4now/Concurrency/race_condition/excercise.go:15 +0xe4
// ==================
// 100
// Found 1 data race(s)
// exit status 66
