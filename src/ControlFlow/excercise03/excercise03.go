package main

import "fmt"

// print out the years you have been alived.

// diff ways of for loop

// func main() {
// 	bd := 1987
// 	for bd <= 2019 {
// 		println(bd)
// 		bd++

// 	}
// }

// func main() {
// 	bd := 1987
// 	for {
// 		if bd > 2019 {
// 			break
// 		}
// 		println(bd)
// 		bd++
// 	}
// }

//=================
// print out the remainder

// func main() {
// 	for i := 10; i <= 100; i++ {
// 		// j := i % 4
// 		// fmt.Printf("%v\n", j)
// 		// or
// 		fmt.Printf("%v\n", i%4)
// 	}
// }

// switch statement

func main() {
	a := "Pyhton"
	switch a {
	case "python":
		fmt.Print(a)
	case "Pyhton":
		fmt.Print(a)
	default:
		print("python") // if the above conditions fails then default will happen
	}

}
