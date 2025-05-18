package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Returns a pointer to Person
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Without pointer - changes won't affect original
func (p Person) BirthdayNonPointer() {
	p.age++ // Only affects local copy
}

// With pointer - changes affect original
func (p *Person) Birthday() {
	p.age++ // Affects original Person
}

// Example of defer
func ShowDeferExample() {
	fmt.Println("Starting function")
	defer fmt.Println("This runs last (3)")
	defer fmt.Println("This runs second (2)")
	fmt.Println("This runs first (1)")
}

func main() {
	// Pointer example
	person := NewPerson("Alice", 20)
	fmt.Printf("Initial age: %d\n", person.age)

	person.BirthdayNonPointer()
	fmt.Printf("After non-pointer birthday: %d\n", person.age)

	person.Birthday()
	fmt.Printf("After pointer birthday: %d\n", person.age)

	// Defer example
	fmt.Println("\nDefer example:")
	ShowDeferExample()
}
