package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Inside \"foo\"")
}

func bar() {
	fmt.Println("Inside \"bar\"")
	defer foo()
	fmt.Println("Call to foo was deferred")
}
