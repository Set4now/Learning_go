package main

import "fmt"

type person struct {
	name string
}

func (r person) speak() string {
	return fmt.Sprintln("I as type person can speak", r.name)
}

type human interface {
	speak() string
}

func saySomething(h human) string {
	return h.speak()
}

func main() {

	p1 := person{
		name: "Suman Deb",
	}

	println(saySomething(&p1))
}
