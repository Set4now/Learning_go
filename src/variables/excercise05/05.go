package main

import "fmt"

// creating custom TYPE
type mytype int

var x mytype
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
	y := mytype(100)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}
