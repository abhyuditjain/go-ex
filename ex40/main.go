package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

// Attached this function to all values of type "rectangle"
func (r rectangle) area() float64 {
	return r.length * r.width
}

// Attached this function to all values of type "circle"
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Anything that has "func rea() float64" attached is also of type "shape"
type shape interface {
	area() float64
}

// polymorphism, "info" can accept anything of type "shape".
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := rectangle{
		length: 2,
		width:  2,
	}

	c := circle{
		radius: 2,
	}

	info(s)
	info(c)
}
