package tickertimer

import (
	"fmt"
	"time"
)

func Example() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("Ticker stopped")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}
