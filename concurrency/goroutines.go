package concurrency

import (
	"fmt"
	"time"
)

// SimpleGoroutine demonstrates a basic goroutine
func SimpleGoroutine() {
	// Launch a goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("From goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Main function continues execution
	for i := 0; i < 3; i++ {
		fmt.Println("From main function:", i)
		time.Sleep(200 * time.Millisecond)
	}

	// Sleep to allow goroutine to finish
	// In real applications, you would use proper synchronization
	time.Sleep(1 * time.Second)
}

// ChannelCommunication demonstrates communication between goroutines using channels
func ChannelCommunication() {
	// Create a channel
	messages := make(chan string)

	// Launch a goroutine that sends a message
	go func() {
		fmt.Println("Sending message from goroutine")
		messages <- "Hello from goroutine!"
	}()

	// Receive the message in the main function
	msg := <-messages
	fmt.Println("Received:", msg)
}
