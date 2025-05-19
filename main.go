package main

import (
	"fmt"
	"github.com/chrubrar/go-learning/concurrency"
)

func main() {
	fmt.Println("=== Demonstrating WaitGroup ===\n")
	concurrency.DemonstrateWaitGroup()

	fmt.Println("\n=== Demonstrating Mutex ===\n")
	concurrency.DemonstrateMutex()

	fmt.Println("\n=== Demonstrating sync.Once ===\n")
	concurrency.DemonstrateSyncOnce()


}
