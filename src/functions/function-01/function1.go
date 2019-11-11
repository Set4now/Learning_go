package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

// Returning a int
func foo() int {
	return 5
}

// Returning both int and string
func bar() (int, string) {
	return 100, "This is a string"
}
