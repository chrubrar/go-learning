package main

import "fmt"

func main() {
	// Regular variable
	num := 42
	
	// & operator: gets the memory address of num
	pointer := &num
	
	fmt.Printf("num's value: %d\n", num)
	fmt.Printf("num's memory address (&num): %p\n", &num)
	fmt.Printf("pointer's value (same as &num): %p\n", pointer)
	fmt.Printf("value at pointer's address (*pointer): %d\n", *pointer)
	
	// Modifying value through pointer
	*pointer = 100
	fmt.Printf("\nAfter *pointer = 100:\n")
	fmt.Printf("num's new value: %d\n", num)
	
	// Example with a function
	fmt.Printf("\nCalling function that takes a pointer:\n")
	modifyValue(&num)
	fmt.Printf("num's value after function call: %d\n", num)
}

func modifyValue(ptr *int) {
	// *ptr accesses the value at the pointer's address
	*ptr = *ptr * 2
	fmt.Printf("Inside function, doubled value: %d\n", *ptr)
}
