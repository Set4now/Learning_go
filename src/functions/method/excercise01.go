package main

import "fmt"

// Creating a Struct
type person struct {
	first string
	last  string
	age   int
}

// Creating a method speak on type person with the receiver format.
func (r person) speak() string {
	// var final_string string
	// string = r.first + r.first + "" + r.age
	return fmt.Sprintf("Info : %v%v Age:- %v", r.first, r.last, r.age)
}

func main() {
	p1 := person{
		first: "Suman",
		last:  "Deb",
		age:   31,
	}
	// Caling the method on a value of type person
	println(p1.speak())
}

// Output
// Info : SumanDeb Age:- 31
