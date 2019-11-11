package main

import "fmt"

// const (
// 	_ = iota
// 	a = iota
// 	b
// 	c

// )

// More examples https://github.com/golang/go/wiki/Iota
const (
	_ = iota // throw way first value of iota
	// << is shift operator so here we are shifing 1 by the expression value (iota * 10)
	kb = 1 << (iota * 10) // iota is incremented by 1 ang shift by 10 bits
	mb = 1 << (iota * 10) // by 20 bits
	gb = 1 << (iota * 10) // by 30 bits
)

func main() {
	// println(a)
	// println(b)
	// println(c)
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
