package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred."},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, slice := range x {
		fmt.Printf("%v\n", i)
		for j, element := range slice {
			fmt.Printf("\t%v\t%v\n", j, element)
		}
	}
}
