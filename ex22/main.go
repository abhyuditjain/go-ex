package main

import "fmt"

func main() {
	// array
	m := [5]int{1, 2, 3, 4, 5}

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", m)
}
