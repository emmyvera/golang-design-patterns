package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Worker:", id, "Started job", j, "...")
		time.Sleep(2 * time.Second)
		fmt.Println("Worker:", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const NUM_JOBS = 5

	jobs := make(chan int, NUM_JOBS)
	results := make(chan int, NUM_JOBS)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= NUM_JOBS; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= NUM_JOBS; a++ {
		<-results
	}
}
