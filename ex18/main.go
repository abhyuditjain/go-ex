package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "Bond" {
		fmt.Println("Bond")
	} else if x == "James Bond" {
		fmt.Println("James Bond")
	} else {
		fmt.Println("Neither")
	}
}
