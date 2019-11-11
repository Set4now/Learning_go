package main

import (
	"fmt"

	"github.com/Set4now/variables/excercise03"
)

func main() {
	// := is short declaration operator only works within func
	X := 42
	Y := "James Bond"
	Z := true

	fmt.Print(X, Y, Z)   // print without new lines
	fmt.Println(X, Y, Z) // print with new lines
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	// This is getting imported from excercise03 package
	fmt.Println(excercise03.Suman())

}
