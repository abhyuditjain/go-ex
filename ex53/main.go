package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// Create waitgroup to sync goroutines, to let them finish work
	var wg sync.WaitGroup

	// Add number of goroutines to waitgroup
	// If this number is less than the goroutines, then program will Panic as wg.Done() will reduce the number in waitgroup by one each time and the number can't be negative
	// If this number is more than the goroutines, then program will keep wating as the number of wg.Done() calls won't make the number 0 and wg.Wait() will keep waiting
	wg.Add(2)

	go func() {
		fmt.Println("Hello")
		// Mark this as done
		// This will reduce the wg count by one
		wg.Done()
	}()

	go func() {
		fmt.Println("World")
		// Mark this as done
		// This will reduce the wg count by one
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())

	// Wait for goroutines to finish
	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
}
