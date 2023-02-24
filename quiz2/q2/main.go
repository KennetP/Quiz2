package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2

	}

}
func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)

	}
	//send 5 jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	//collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println(<-results)
	}
}
