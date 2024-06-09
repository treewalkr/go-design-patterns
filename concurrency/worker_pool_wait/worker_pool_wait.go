package workerpoolwait

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
	}
}

func Example() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers
	wg.Wait()
	fmt.Println("All workers completed")
}
