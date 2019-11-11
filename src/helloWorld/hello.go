package main

import "fmt"

// Every script should have this package main and main funtion that is entrypoint
// All other code is sequential inside the main function
var b int

func main() {
	fmt.Printf("hello, world\n")
	a := 20
	b = 100
	fmt.Print(a)
	fmt.Println(b)
	nextfuc()
}

func nextfuc() {
	fmt.Println("I am exiting")
}
