package main

import "fmt"

func main() {
	// slice
	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", m)

	// : is slicing operator
	// left side is inclusive
	// right is exclusive
	fmt.Println(m[:5])
	fmt.Println(m[5:])
	fmt.Println(m[2:7])
	fmt.Println(m[1:6])
}
