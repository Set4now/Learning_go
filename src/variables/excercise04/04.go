package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
