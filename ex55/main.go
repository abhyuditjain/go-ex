package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// the aim is to create a race condition
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	iterations := 100

	var wg sync.WaitGroup
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			// read value
			local := counter
			fmt.Println("Local", local)
			// yield the processor
			runtime.Gosched()

			local++

			counter = local

			fmt.Println("Counter written:", counter)

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter:", counter)
}
