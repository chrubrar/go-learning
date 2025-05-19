package examples

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// ==========================================
// SECTION 1: BASIC SYNTAX AND DATA TYPES
// ==========================================

func basicTypes() {
	// Variable declarations
	var a int = 10
	var b = 20       // Type inferred
	c := 30          // Short declaration
	
	fmt.Println("Basic Types:")
	fmt.Println("Integers:", a, b, c)
	
	// Multiple declarations
	var (
		name   string = "John"
		age    int    = 30
		salary float64 = 50000.50
	)
	fmt.Println("Multiple variables:", name, age, salary)
	
	// Constants
	const Pi = 3.14159
	fmt.Println("Constant Pi:", Pi)
	
	// Basic types
	var (
		intVar    int     = 42
		int32Var  int32   = 42
		int64Var  int64   = 42
		uintVar   uint    = 42
		floatVar  float64 = 42.42
		boolVar   bool    = true
		stringVar string  = "Hello, Go!"
		runeVar   rune    = 'A'
		byteVar   byte    = 65
	)
	
	fmt.Println("Integer:", intVar)
	fmt.Println("Int32:", int32Var)
	fmt.Println("Int64:", int64Var)
	fmt.Println("Unsigned int:", uintVar)
	fmt.Println("Float64:", floatVar)
	fmt.Println("Boolean:", boolVar)
	fmt.Println("String:", stringVar)
	fmt.Println("Rune (Unicode code point):", runeVar, string(runeVar))
	fmt.Println("Byte:", byteVar, string(byteVar))
}

// ==========================================
// SECTION 2: CONTROL STRUCTURES
// ==========================================

func controlStructures() {
	fmt.Println("\nControl Structures:")
	
	// If-else statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is equal to 5")
	}
	
	// If with a short statement
	if y := 15; y > 10 {
		fmt.Println("y is greater than 10")
	}
	
	// For loop (standard C-style)
	fmt.Print("Standard for loop: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// For loop as a while loop
	fmt.Print("For as while: ")
	j := 0
	for j < 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
	
	// Infinite loop with break
	fmt.Print("Loop with break: ")
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Print(k, " ")
		k++
	}
	fmt.Println()
	
	// Loop with continue
	fmt.Print("Loop with continue (skip even numbers): ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// Switch statement
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	case "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's the weekend")
	}
	
	// Switch without an expression (like if-else)
	hour := 15
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// ==========================================
// SECTION 3: ARRAYS, SLICES, AND MAPS
// ==========================================

func collectionsDemo() {
	fmt.Println("\nArrays, Slices, and Maps:")
	
	// Arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	
	// Array initialization
	arr2 := [3]string{"apple", "banana", "cherry"}
	fmt.Println("String array:", arr2)
	
	// Array with size determined by initialization
	arr3 := [...]int{10, 20, 30, 40}
	fmt.Println("Array with inferred size:", arr3, "Length:", len(arr3))
	
	// Slices (dynamic arrays)
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice1, "Length:", len(slice1), "Capacity:", cap(slice1))
	
	// Creating a slice from an array
	sliceFromArray := arr[1:4]  // elements 1, 2, 3 (indices 1, 2, 3)
	fmt.Println("Slice from array:", sliceFromArray)
	
	// Creating a slice with make
	slice2 := make([]int, 5, 10)  // length 5, capacity 10
	fmt.Println("Slice created with make:", slice2, "Length:", len(slice2), "Capacity:", cap(slice2))
	
	// Appending to a slice
	slice1 = append(slice1, 6, 7)
	fmt.Println("After append:", slice1)
	
	// Appending one slice to another
	slice3 := []int{8, 9, 10}
	slice1 = append(slice1, slice3...)  // ... is the spread operator
	fmt.Println("After appending another slice:", slice1)
	
	// Maps (key-value pairs, like dictionaries or hash tables)
	map1 := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 12,
	}
	fmt.Println("Map:", map1)
	
	// Creating a map with make
	map2 := make(map[string]float64)
	map2["pi"] = 3.14159
	map2["e"] = 2.71828
	fmt.Println("Map created with make:", map2)
	
	// Checking if a key exists
	value, exists := map1["apple"]
	if exists {
		fmt.Println("apple exists in the map with value:", value)
	}
	
	// Deleting from a map
	delete(map1, "banana")
	fmt.Println("Map after delete:", map1)
	
	// Iterating over a map
	fmt.Println("Map iteration:")
	for key, value := range map2 {
		fmt.Printf("  %s: %f\n", key, value)
	}
}

// ==========================================
// SECTION 4: FUNCTIONS
// ==========================================

// Basic function
func add(a, b int) int {
	return a + b
}

// Multiple return values
func divideAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// Named return values
func calculateStats(numbers []int) (min, max, total int) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	total = 0
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		total += num
	}
	
	return  // "naked" return - returns the named return values
}

// Variadic function (variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as a value (first-class functions)
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Closure (function that references variables from outside its body)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func functionsDemo() {
	fmt.Println("\nFunctions:")
	
	// Basic function call
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
	
	// Multiple return values
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("10 / 3 =", quotient, "with remainder", remainder)
	
	// Named return values
	numbers := []int{7, 2, 9, 4, 13}
	min, max, sum := calculateStats(numbers)
	fmt.Printf("For %v: min=%d, max=%d, sum=%d\n", numbers, min, max, sum)
	
	// Variadic function
	//fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))
	
	// Passing a slice to a variadic function
	//nums := []int{10, 20, 30, 40}
	//fmt.Println("Sum of", nums, ":", sum(nums...))
	
	// Function as a value
	multiply := func(x, y int) int { return x * y }
	fmt.Println("5 * 3 =", applyOperation(5, 3, multiply))
	
	// Anonymous function
	fmt.Println("5 - 3 =", applyOperation(5, 3, func(x, y int) int { return x - y }))
	
	// Closure
	counter := makeCounter()
	fmt.Println("Counter:", counter())  // 1
	fmt.Println("Counter:", counter())  // 2
	fmt.Println("Counter:", counter())  // 3
	
	// Each closure has its own state
	counter2 := makeCounter()
	fmt.Println("Counter2:", counter2())  // 1
	fmt.Println("Counter:", counter())    // 4
}

// ==========================================
// SECTION 5: STRUCTS AND METHODS
// ==========================================

// Define a struct (similar to a class in other languages)
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with a receiver (like a class method)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Nested structs
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

// Embedded structs (composition)
type Employee struct {
	Person  // Embedded struct
	JobTitle string
	Salary   float64
}

func structsDemo() {
	fmt.Println("\nStructs and Methods:")
	
	// Create a struct instance
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Rectangle:", rect)
	
	// Access struct fields
	fmt.Println("Width:", rect.Width, "Height:", rect.Height)
	
	// Call a method
	fmt.Println("Area:", rect.Area())
	
	// Call a method with a pointer receiver
	rect.Scale(2)
	fmt.Println("After scaling:", rect)
	fmt.Println("New area:", rect.Area())
	
	// Nested structs
	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
	}
	fmt.Println("Person:", person)
	fmt.Println("Address:", person.Address.Street, person.Address.City)
	
	// Embedded structs
	employee := Employee{
		Person: Person{
			Name: "Jane Smith",
			Age:  28,
			Address: Address{
				Street:  "456 Oak Ave",
				City:    "Somewhere",
				Country: "Canada",
			},
		},
		JobTitle: "Software Engineer",
		Salary:   85000,
	}
	fmt.Println("Employee:", employee)
	
	// With embedded structs, fields can be accessed directly
	fmt.Println("Employee name:", employee.Name)  // Instead of employee.Person.Name
	fmt.Println("Employee city:", employee.Address.City)
}

// ==========================================
// SECTION 6: INTERFACES
// ==========================================

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle already has Area(), let's add Perimeter()
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function that accepts an interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Empty interface can hold any type
func describeValue(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func interfacesDemo() {
	fmt.Println("\nInterfaces:")
	
	// Create shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}
	
	// Use the interface
	fmt.Print("Circle: ")
	printShapeInfo(circle)
	
	fmt.Print("Rectangle: ")
	printShapeInfo(rectangle)
	
	// Slice of interfaces
	shapes := []Shape{circle, rectangle}
	fmt.Println("All shapes:")
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
	
	// Type assertions
	var s Shape = Circle{Radius: 3}
	
	// Check if s is a Circle
	if c, ok := s.(Circle); ok {
		fmt.Println("s is a Circle with radius", c.Radius)
	}
	
	// Type switch
	switch v := s.(type) {
	case Circle:
		fmt.Println("Circle with radius", v.Radius)
	case Rectangle:
		fmt.Println("Rectangle with width", v.Width, "and height", v.Height)
	default:
		fmt.Println("Unknown shape")
	}
	
	// Empty interface
	fmt.Println("Empty interface examples:")
	describeValue(42)
	describeValue("hello")
	describeValue(true)
	describeValue(circle)
}

// ==========================================
// SECTION 7: CONCURRENCY WITH GOROUTINES AND CHANNELS
// ==========================================

func concurrencyDemo() {
	fmt.Println("\nConcurrency with Goroutines and Channels:")
	
	// Basic goroutine
	go func() {
		fmt.Println("Hello from a goroutine!")
	}()
	
	// Give the goroutine time to execute
	time.Sleep(100 * time.Millisecond)
	
	// Channel basics
	fmt.Println("Channel basics:")
	ch := make(chan string)
	
	go func() {
		ch <- "Message from goroutine"
	}()
	
	msg := <-ch
	fmt.Println("Received:", msg)
	
	// Buffered channels
	fmt.Println("Buffered channels:")
	bufferedCh := make(chan int, 3)
	
	// Send values without a receiver
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	// Receive values
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	
	// Channel directions (send-only and receive-only)
	fmt.Println("Channel directions:")
	
	// Function that only sends to a channel
	sender := func(ch chan<- string) {
		ch <- "Hello, receiver!"
	}
	
	// Function that only receives from a channel
	receiver := func(ch <-chan string) {
		msg := <-ch
		fmt.Println("Receiver got:", msg)
	}
	
	ch2 := make(chan string)
	go sender(ch2)
	go receiver(ch2)
	
	time.Sleep(100 * time.Millisecond)
	
	// Select statement
	fmt.Println("Select statement:")
	ch3 := make(chan string)
	ch4 := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch3 <- "Message from ch3"
	}()
	
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch4 <- "Message from ch4"
	}()
	
	// Select waits for either channel to be ready
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch3:
			fmt.Println("Received from ch3:", msg)
		case msg := <-ch4:
			fmt.Println("Received from ch4:", msg)
		}
	}
	
	// WaitGroup for synchronization
	fmt.Println("WaitGroup for synchronization:")
	var wg sync.WaitGroup
	
	// Launch 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)  // Increment counter
		
		go func(id int) {
			defer wg.Done()  // Decrement counter when done
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}
	
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines completed")
	
	// Mutex for shared memory synchronization
	fmt.Println("Mutex for shared memory synchronization:")
	var (
		counter = 0
		mutex   sync.Mutex
	)
	
	increment := func() {
		mutex.Lock()
		defer mutex.Unlock()
		counter++
	}
	
	// Launch 1000 goroutines that increment the counter
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			increment()
		}()
	}
	
	wg2.Wait()
	fmt.Println("Counter value:", counter)
}

// ==========================================
// SECTION 8: ERROR HANDLING
// ==========================================

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Custom error type
type ValidationError struct {
	Field string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

// Function that returns a custom error
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{Field: "age", Message: "age cannot be negative"}
	}
	if age > 150 {
		return ValidationError{Field: "age", Message: "age is unrealistically high"}
	}
	return nil
}

func errorHandlingDemo() {
	fmt.Println("\nError Handling:")
	
	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}
	
	// Custom error
	err = validateAge(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age 25 is valid")
	}
	
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Error:", err)
		
		// Type assertion to check for specific error type
		if validationErr, ok := err.(ValidationError); ok {
			fmt.Println("Validation error details:")
			fmt.Println("  Field:", validationErr.Field)
			fmt.Println("  Message:", validationErr.Message)
		}
	}
	
	// Panic and recover
	fmt.Println("Panic and recover:")
	
	// Function that panics
	panicFunc := func() {
		fmt.Println("panicFunc: start")
		panic("something went wrong")
		fmt.Println("panicFunc: end")  // This won't be executed
	}
	
	// Function that recovers from panic
	recoverFunc := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		
		fmt.Println("recoverFunc: calling panicFunc")
		panicFunc()
		fmt.Println("recoverFunc: after panicFunc")  // This won't be executed
	}
	
	fmt.Println("main: calling recoverFunc")
	recoverFunc()
	fmt.Println("main: after recoverFunc")
}

// ==========================================
// MAIN FUNCTION
// ==========================================

func main() {
	fmt.Println("========== GO PROGRAMMING TUTORIAL ==========")
	
	// Uncomment the sections you want to run
	basicTypes()
	controlStructures()
	collectionsDemo()
	//functionsDemo()
	structsDemo()
	interfacesDemo()
	concurrencyDemo()
	errorHandlingDemo()
	
	fmt.Println("\n========== END OF TUTORIAL ==========")
}
