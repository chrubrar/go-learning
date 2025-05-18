package main

import (
	"fmt"

	"github.com/chrubrar/go-learning/basics"
	"github.com/chrubrar/go-learning/concurrency"
)

func main() {
	fmt.Println("Hello, Go World!")

	// Using the Max function from basics package
	result := basics.Max(10, 5)
	fmt.Println("Max of 10 and 5 is:", result)

	// Using goroutines from concurrency package
	fmt.Println("\nDemonstrating simple goroutine:")
	concurrency.SimpleGoroutine()

	fmt.Println("\nDemonstrating channel communication:")
	concurrency.ChannelCommunication()
}
