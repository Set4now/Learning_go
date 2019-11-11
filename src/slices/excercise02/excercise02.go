package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 44, 45, 46, 47, 48, 49, 50, 51} // creating a slice
	x = append(x, 52)                                      // appending an item to a slice
	fmt.Printf("%v\n", x)
	x = append(x, 53, 54, 55) // appending multiple items to the slice
	fmt.Printf("%v\n", x)
	y := []int{56, 57, 58, 59, 60} // creating another slice
	x = append(x, y...)            // appending the slice y to slice x | ... is important
	fmt.Printf("%v\n", x)
}

// Output:-
// [42 43 44 44 45 46 47 48 49 50 51 52]
// [42 43 44 44 45 46 47 48 49 50 51 52 53 54 55]
// [42 43 44 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
