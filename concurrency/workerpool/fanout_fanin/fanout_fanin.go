package fanoutfanin

import "fmt"

func producer(nums []int, out chan<- int) {
	for _, num := range nums {
		out <- num
	}
	close(out)
}

func worker(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
}

func Example() {
	nums := []int{1, 2, 3, 4, 5}
	in := make(chan int, len(nums))
	out := make(chan int, len(nums))

	go producer(nums, in)

	// Fan-Out: Start multiple workers
	for i := 0; i < 3; i++ {
		go worker(in, out)
	}

	// Fan-In: Collect results
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-out)
	}
}
