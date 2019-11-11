package main

import (
	"runtime"
	"sync"
)

// create a variable of type Waitgroup from sync package
var gr sync.WaitGroup

func main() {
	// Printing Number of CPU's and Goroutines from packge runtime
	println("Number of CPU's ", runtime.NumCPU())
	println("GoRoutines ", runtime.NumGoroutine())
	students := make([]string, 0)

	students = []string{"suman", "Sunetra"}
	// Adding Delta which is the required number of Goroutines.
	// So we are dynamically assiging it by the length of our students variab;e
	gr.Add(len(students))
	for _, v := range students {
		go student(v)

		// while i work schedule another goroutine
		// time.Sleep(2 * time.Second)
		// runtime.Gosched()

	}
	println("Number of GoRoutines at work ", runtime.NumGoroutine())
	// The main function will wait till the time Go routnines execution is over.
	gr.Wait()
	println("GoRoutines at the end ", runtime.NumGoroutine())
}

func student(name string) {
	println("", name)
	// My go routine execution is done.
	gr.Done()
}
