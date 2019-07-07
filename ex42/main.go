package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// anonymous function assigned to "total"
	total := func(i ...int) int {
		total := 0
		for _, v := range i {
			total += v
		}
		return total
	}

	// function called
	fmt.Println(total(x...))
}
