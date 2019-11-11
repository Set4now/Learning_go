package main

import "fmt"

func main() {

	type person struct {
		first            string
		last             string
		icecreamflavours []string
	}

	p1 := person{
		first:            "Suman",
		last:             "Deb",
		icecreamflavours: []string{"vanilla", "chocolate"},
	}

	fmt.Println(p1)
	fmt.Println("Printing the icecream flavours:- ")
	for _, v := range p1.icecreamflavours {
		println(v)
	}
}
