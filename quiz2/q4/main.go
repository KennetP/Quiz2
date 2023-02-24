package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) //increment the waitgroup counter before starting a new waitgroup routine
		go func(id int) {
			defer wg.Done() //decrement the waitgroup counter when go routine is finished
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d finished\n", id)

		}(i)
	}
	wg.Wait() //block main goroutine until all goroutines are finished
	fmt.Println("All Goroutines finished")
}
