package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Attached this function to all values of type "person"
func (s person) speak() {
	fmt.Printf("My name is %s %s and I am %d years old\n", s.first, s.last, s.age)
}

func main() {
	p1 := person{
		first: "Abhyudit",
		last:  "Jain",
		age:   26,
	}

	// Attached functions callable using dot notation
	p1.speak()
}
