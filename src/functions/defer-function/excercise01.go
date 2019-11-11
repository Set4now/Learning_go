package main

import "fmt"

// The defer keyword tells the function to run once the main function enclosing it exits.
func main() {
	defer fmt.Println(mydefer())
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() string {
	return "I am the main function"
}

func bar() string {
	return "I am the main function part 2"
}

func mydefer() string {
	return "I am the defer function"
}
