package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("I am", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Abhyudit",
		last:  "Jain",
	}

	// Works
	saySomething(&p1)
	// Doesn't work
	// saySomething(p1)

	(&p1).speak()
	// Same as above
	p1.speak()
}
