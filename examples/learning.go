package examples

import (
	"fmt"
	"time"
)

// Max returns the larger of two integers
func Max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

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

func BasicForLoop() {
	for x := 1; x <= 10; x++ {
		fmt.Println("We are so happy to see you:", x)
	}
}

func RangeSlices() {
	fmt.Println("\n :Printing range values \n")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)
}

// ExplainGoroutines demonstrates how goroutines work
func ExplainGoroutines() {
	fmt.Println("Starting goroutine demonstration...")

	// Start a goroutine that counts up
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Goroutine counting up: %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// At the same time, main thread counts down
	for i := 5; i >= 1; i-- {
		fmt.Printf("Main thread counting down: %d\n", i)
		time.Sleep(300 * time.Millisecond)
	}

	// Wait for goroutine to finish
	time.Sleep(3 * time.Second)
	fmt.Println("Goroutine demonstration complete")
}

func GetAgeInput() {

	var age int
election: // this is the label

	fmt.Println("Enter your age")

	fmt.Scanln(&age) //user enters age

	if age <= 17 {

		fmt.Println("you are not eligible to vote")
		goto election
	} else {
		fmt.Println("You are eligible to vote")
	}

}

func add(x int, y int) {
	total := x + y
	fmt.Println(total)
}

func Greeting(name string) (message string) {
	message = "hello," + name
	return
}

// DemonstrateAnonymousFunctions shows different ways to use anonymous functions
func DemonstrateAnonymousFunctions() {
	fmt.Println("\n=== Anonymous Functions Demo ===")

	// 1. Immediate execution
	func() {
		fmt.Println("Anonymous function executed immediately")
	}()

	// 2. Assigned to variable
	greet := func(name string) string {
		return fmt.Sprintf("Hello, %s!", name)
	}
	fmt.Println(greet("Gopher"))

	// 3. With closure (accessing outside variables)
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println("Counter:", increment())
	fmt.Println("Counter:", increment())

	// 4. As function parameter
	numbers := []int{1, 2, 3, 4, 5}
	result := processNumbers(numbers, func(x int) int {
		return x * x // square the number
	})
	fmt.Println("Squared numbers:", result)
}

// Helper function that uses an anonymous function as parameter
func processNumbers(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

func main() {
	fmt.Println("Hello, Go World!")

	// Add the new goroutine demonstration
	fmt.Println("\nDemonstrating goroutines in detail:")
	ExplainGoroutines()

	// Using the Max function
	result := Max(10, 5)
	fmt.Println("Max of 10 and 5 is:", result)

	// Using goroutines
	fmt.Println("\nDemonstrating simple goroutine:")
	SimpleGoroutine()

	fmt.Println("\nDemonstrating channel communication:")
	ChannelCommunication()

	//BasicForLoop()

	//RangeSlices()

	//ExplainGoroutines()

	GetAgeInput()

	add(3, 2)

	msg := Greeting("Brooklyn")
	fmt.Println(msg)

	defer fmt.Println("world!")
	fmt.Println("Hello,")

	for a := 1; a <= 3; a++ {
		defer fmt.Println(a)
	}

	fmt.Println("world")

	DemonstrateAnonymousFunctions()
}
