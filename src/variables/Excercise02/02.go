package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	a := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	fmt.Print(a)
}
