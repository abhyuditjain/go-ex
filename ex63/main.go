package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// If I commnet this goroutine, the program panics
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}
