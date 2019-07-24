package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// anonymous self executing funtion
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// buffered channel
	c = make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
