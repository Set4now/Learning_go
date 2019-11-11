package main

// Using Recusrion
// func fact(n int) int {
// 	if n == 1 {
// 		return n
// 	}
// 	return n * fact(n-1)

// }

// Witout Recusrsion
func main() {
	println(myfact(6))
}

func myfact(n int) int {
	last := 1
	for i := 1; i <= n; i++ {
		last = last * i
	}
	return last
}
