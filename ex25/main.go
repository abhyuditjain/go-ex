package main

import "fmt"

func main() {
	// slice
	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	m = append(m, 52)
	fmt.Println(m)
	m = append(m, 53, 54, 55)
	fmt.Println(m)
	m = append(m, []int{56, 57, 58, 59, 60}...)
	fmt.Println(m)
}
