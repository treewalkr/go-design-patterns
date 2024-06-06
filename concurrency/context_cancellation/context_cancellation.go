package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d working\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main function done")
}
