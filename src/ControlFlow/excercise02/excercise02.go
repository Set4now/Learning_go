package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {

		fmt.Printf("\t\t%#U\n", i)
		fmt.Printf("\t\t%#U\n", i)
		fmt.Printf("\t\t%#U\n", i)

	}
}

// Output:-
// m-c02tm0amg8wn:ControlFlow s0d011k$ go run excercise02/excercise02.go
//                 U+0041 'A'
//                 U+0041 'A'
//                 U+0041 'A'
//                 U+0042 'B'
//                 U+0042 'B'
//                 U+0042 'B'
//                 U+0043 'C'
//                 U+0043 'C'
//                 U+0043 'C'
//                 U+0044 'D'
//                 U+0044 'D'
//                 U+0044 'D'
//                 U+0045 'E'

// func main() {

// 	for i := 65; i <= 90; i++ {
// 		fmt.Println(i)
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("\t\t%#U\n", i)

// 		}

// 	}

// }

// Output:-
// m-c02tm0amg8wn:ControlFlow s0d011k$ go run excercise02/excercise02.go
// 65
//                 U+0041 'A'
//                 U+0041 'A'
//                 U+0041 'A'
// 66
//                 U+0042 'B'
//                 U+0042 'B'
//                 U+0042 'B'
// 67
//                 U+0043 'C'
//                 U+0043 'C'
//                 U+0043 'C'
