package main

import "fmt"

func main() {
	// slice
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", m)
}
