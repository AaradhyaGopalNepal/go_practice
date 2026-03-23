package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d starting \n", id)
	time.Sleep(time.Second)
	results <- id * 2
	fmt.Printf("WorkerID %d finished \n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 3
	results := make(chan int, numJobs)
	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range <-results {
		fmt.Println("Result:", result)
	}

}
