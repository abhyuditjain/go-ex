package main

import "fmt"

// Package level scope
// As they are uninitialized, they get assigned ZERO VALUE
// for their respective types
// 0 for int
// "" for string
// false for bool
var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
