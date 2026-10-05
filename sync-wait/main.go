package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal the goroutine is done
	fmt.Printf("Worker %d started\n", i)
	// some tasks is happening
	fmt.Printf("Worker %d ended\n", i)
}
func main() {
	// fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	// start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment the waitgroup timer
		go worker(i, &wg)
	}

	// wait for all the workers to finish
	wg.Wait()

	fmt.Println("Worker tasks ended")
}
