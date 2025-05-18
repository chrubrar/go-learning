package main

import (
	"fmt"
)

func FactorialCalc(number int) int {

	if number == 0 || number == 1 {

		return 1
	}

	if number < 0 {

		fmt.Println("Invalid Number")

		return -1
	}

	return number * FactorialCalc(number-1) // recursive function unil operations reached the base case

}

func print_one(n int) {
	if n > 0 {
		fmt.Println("in first function:", n)
		print_two(n - 1)
	}
}
func print_two(n int) {
	if n >= 0 {
		fmt.Println("In second function:", n)
		print_one(n - 1)
	}
}

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

func processNumbers(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

func sum(a, b int) int {
	return a + b
}

func partialSum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {

	answer1 := FactorialCalc(0)

	fmt.Println(answer1, "\n")

	answer2 := FactorialCalc(4)

	fmt.Println(answer2, "\n")

	answer3 := FactorialCalc(7)

	fmt.Println(answer3)

	print_one(6)

	print_one(-10)

	DemonstrateAnonymousFunctions()
	partial := partialSum(2)
	fmt.Println(partial(5))
	fmt.Println(partial(12))

}
