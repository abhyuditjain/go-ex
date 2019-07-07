package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Double:", mapper(double, x))
	fmt.Println("Triple:", mapper(triple, x))
}

func mapper(f func(i int) int, x []int) []int {
	mapped := make([]int, len(x), cap(x))
	for i, v := range x {
		mapped[i] = f(v)
	}
	return mapped
}

func double(x int) int {
	return 2 * x
}

func triple(x int) int {
	return 3 * x
}
