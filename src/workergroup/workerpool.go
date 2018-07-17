package main

import (
	"fmt"
)

func squareNumberWorker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Create a few workers
	for i := 1; i < 10; i++ {
		go squareNumberWorker(jobs, results)
	}

	// Make 50 jobs
	for j := 1; j < 100; j++ {
		jobs <- j
	}
	close(jobs)

	// Grab all results in array
	var squares [100]int
	for i := 1; i < 100; i++ {
		squares[i] = <-results
	}
	close(results)
	fmt.Println(squares)
}
