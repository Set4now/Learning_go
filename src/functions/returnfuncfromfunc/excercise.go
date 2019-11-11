package main

import "fmt"

func main() {
	a := test()
	fmt.Println(a())
}

func test() func() string {
	return func() string {
		return "I will be returned inside another function"
	}

}
