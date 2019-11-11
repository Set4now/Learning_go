package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // creating a slice of 10 values
	// If you want to print the index along with its values then replace _ with any identifier and print that along with v
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T", x)

}

// output:=
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// []int----> type of the underlying array which slice creates.
