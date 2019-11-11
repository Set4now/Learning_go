package main

import "fmt"

func main() {
	a := "suman"
	b := "Sunetra"
	c := "swati"
	d := "leena"

	slice1 := []string{a, b, c, d}
	// fmt.Printf("%v\n", slice1)
	slice2 := []string{"Engineer", "Mom", "Wife", "Sister"}
	// fmt.Printf("%v\n", slice2)
	finalslice := [][]string{slice1, slice2}
	// fmt.Printf("%v\n", finalslice)
	// Iteration over items on main slice and then on each item
	// for k := 0; k < len(finalslice); k++ {

	for l, v := range finalslice {

		// for k := 0; k < len(finalslice); k++ {
		fmt.Printf("%v\t\n", l)
		for i, j := range v {

			fmt.Printf("\t\t%v\t%v\n", i, j)

		}

	}

}

// Output
// 0
//                 0       suman
//                 1       Sunetra
//                 2       swati
//                 3       leena
// 1
//                 0       Engineer
//                 1       Mom
//                 2       Wife
//                 3       Sister
