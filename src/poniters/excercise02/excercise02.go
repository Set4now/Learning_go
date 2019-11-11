package main

import "fmt"

// Creating a struct
type person struct {
	name string
}

func main() {
	p := person{
		name: "Suman Deb",
	}

	fmt.Printf("%v\n", p)
	// Calling changeMe and passing addess of p using & operator which is nothing but type of " * person"
	changeMe(&p)
	fmt.Printf("%v\n", p)
}

// Creating a func that takes a value of type PONITER to person
func changeMe(p *person) {
	p.name = "Sunetra Das"
}
