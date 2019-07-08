package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)   // value
	fmt.Println(&x)  // address
	fmt.Println(*&x) // value
}
