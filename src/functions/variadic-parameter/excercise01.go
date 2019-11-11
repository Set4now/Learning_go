package main

// Sum of all numbers in a slice
func main() {
	y := []int{10, 13, 17}
	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	println(foo(y...))

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
