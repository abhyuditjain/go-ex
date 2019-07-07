package main

import "fmt"

func main() {
	x := foo()
	y, name := bar()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(name)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 7, "James Bond"
}
