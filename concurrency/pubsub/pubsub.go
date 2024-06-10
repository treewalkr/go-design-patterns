package pubsub

import (
	"fmt"
	"sync"
	"time"
)

// Broker manages the pub/sub system, handling message distribution between publishers and subscribers
type Broker struct {
	subscribers map[chan string]struct{} // Map to store subscriber channels
	mu          sync.Mutex               // Mutex to ensure thread-safe operations
}

// NewBroker creates and initializes a new Broker instance
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[chan string]struct{}),
	}
}

// Subscribe creates a new subscription channel and registers it with the broker.
// Returns a channel that will receive published messages.
func (b *Broker) Subscribe() chan string {
	ch := make(chan string)
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[ch] = struct{}{} // Add channel to subscribers map
	return ch
}

// Publish broadcasts a message to all registered subscribers.
// Each subscriber receives the message through its dedicated channel.
func (b *Broker) Publish(msg string) {
	fmt.Printf("Publishing message: %s\n", msg)
	b.mu.Lock()
	defer b.mu.Unlock()

	// Send message to each subscriber in a non-blocking way
	for ch := range b.subscribers {
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}
}

// Unsubscribe removes a subscription channel from the broker.
// Closes the channel to signal the subscriber to stop listening.
func (b *Broker) Unsubscribe(ch chan string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.subscribers, ch)
	close(ch)
}

// Example demonstrates how to use the pub/sub system with multiple publishers and subscribers
func Example() {
	broker := NewBroker()

	// Set up first subscriber
	sub1 := broker.Subscribe()
	go func() {
		for msg := range sub1 {
			fmt.Println("Subscriber 1 received: ", msg)
		}
	}()

	// Set up second subscriber
	sub2 := broker.Subscribe()
	go func() {
		for msg := range sub2 {
			fmt.Println("Subscriber 2 received: ", msg)
		}
	}()

	// Set up publisher that sends 10 messages with 1-second intervals
	go func() {
		for i := 0; i < 10; i++ {
			broker.Publish(fmt.Sprintf("Message %d", i))
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second) // Wait for messages to be published
}
