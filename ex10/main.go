package main

import "fmt"

func main() {
	// Raw string literal. Will keep all the characters including spaces and newlines
	a := `Hello, 
	
				   world!`

	fmt.Println(a)
}
