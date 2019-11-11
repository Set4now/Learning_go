package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("I have been assigned to a variable")
	}
	a()
}
