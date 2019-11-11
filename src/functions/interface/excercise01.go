package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := square{
		length: 10,
	}
	c1 := circle{
		radius: 50.0,
	}

	fmt.Println(info(s1))
	fmt.Println(info(c1))

}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	// pii := 3.14

	return c.radius * c.radius * math.Pi

}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(sh shape) float64 {
	return sh.area()
}
