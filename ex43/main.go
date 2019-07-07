package main

import "fmt"

func main() {
	x := foo()

	x()
}

func foo() func() {
	fmt.Println("Inside outside func")
	return func() {
		fmt.Println("Inside inside func")
	}
}
