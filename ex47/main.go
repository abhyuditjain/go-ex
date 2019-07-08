package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "AbhyuFit"
	// Same as above
	// (*p).first = "AbhyuFit"
}

func main() {
	p := person{
		first: "Abhyudit",
		last:  "Jain",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
