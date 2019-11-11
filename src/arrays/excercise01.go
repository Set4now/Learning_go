package main

import "fmt"

func main() {
	// var x [5]int //you can declare like this and assign elements to inidivual index position.
	// x[0] = 1
	// x[1] = 2
	// x[2] = 3
	// x[4] = 4
	x := [5]int{1, 2, 3, 4, 5} // More recommended single line 5 defines the size which is the number of
	//index elements this can hold
	// _ is used to ignore a variable
	for _, v := range x {
		fmt.Printf("%v\t%T\n", v, v)
	}
	fmt.Printf("%T", x) //This will give size and type of the array
	// output:-
	// 1       int
	// 2       int
	// 3       int
	// 4       int
	// 5       int
	// [5]	   int ---> size of the array and type of the array
}
