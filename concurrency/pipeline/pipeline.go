package pipeline

import "fmt"

func stage1(input []int, out chan<- int) {
	for _, num := range input {
		out <- num + 1
	}
	close(out)
}

func stage2(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func Example() {
	numbers := []int{1, 2, 3, 4, 5}
	stage1Out := make(chan int, len(numbers))
	stage2Out := make(chan int, len(numbers))

	go stage1(numbers, stage1Out)
	go stage2(stage1Out, stage2Out)

	for result := range stage2Out {
		fmt.Println(result)
	}
}
