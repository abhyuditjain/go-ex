package main

import "fmt"

// Custom type which has underlying type "int"
type mytype int

var x mytype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
