// deleting items from a slice

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// use slicing method and append to create a slice that will look like
	// [42, 43, 44, 48, 49, 50, 51]
	x = append(x[:3], x[6:]...) // x[:3] and x[6:1] is the slciing methods
	fmt.Printf("%v\n", x)

}

// output
// [42 43 44 48 49 50 51]
