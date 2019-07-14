package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// the aim is to fix the race condition using mutex
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	// Create lock
	var lock sync.Mutex

	iterations := 100

	var wg sync.WaitGroup
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			// lock
			lock.Lock()
			// read value
			local := counter
			fmt.Println("Local", local)
			// yield the processor
			// don't need to yield the processor
			// runtime.Gosched()

			local++

			counter = local

			fmt.Println("Counter written:", counter)

			// unlock
			lock.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter:", counter)
}
