package examples

import "fmt"

// DemonstrateClosures shows different ways to use closures
func DemonstrateClosures() {
	fmt.Println("=== Closure Examples ===")

	// Example 1: Simple Counter
	counter := createCounter()
	fmt.Println(counter()) // prints 1
	fmt.Println(counter()) // prints 2
	fmt.Println(counter()) // prints 3

	// Example 2: Counter with initial value
	counterFrom5 := createCounterWithStart(5)
	fmt.Println(counterFrom5()) // prints 6
	fmt.Println(counterFrom5()) // prints 7

	// Example 3: Adder function
	addTen := createAdder(10)
	fmt.Println(addTen(5))  // prints 15
	fmt.Println(addTen(20)) // prints 30

	// Example 4: Message formatter
	formatter := createFormatter("Hello")
	fmt.Println(formatter("John")) // prints "Hello, John!"
	fmt.Println(formatter("Jane")) // prints "Hello, Jane!"

	formatter1 := createFormatter1("Hi")
	fmt.Println(formatter1("taig"))
}

// Example 1: Simple counter closure
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Example 2: Counter with custom start value
func createCounterWithStart(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// Example 3: Adder closure
func createAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

// Example 4: Message formatter closure
func createFormatter(prefix string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", prefix, name)
	}
}

func createFormatter1(str string) func(string) string {
	return func(name1 string) string {
		return fmt.Sprintf("%s, %s!", str, name1)
	}
}

func main() {
	DemonstrateClosures()
}
