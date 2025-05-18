package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
)

func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Add(1, 2) =", Add(1, 2))
	fmt.Println("math.Sqrt(4) =", math.Sqrt(4))
	fmt.Println("rand.Intn(10) =", rand.Intn(10))

	matched, err := regexp.MatchString("foo", "foobar")

	fmt.Println(matched)
	fmt.Println(err)

	matched1, err := regexp.MatchString(`^a.{1,3}b$`, "aaccbb")

	fmt.Println(matched1)
	fmt.Println(err)

	// Compile example
	pattern := `\d{3}-\d{2}-\d{4}` // Pattern for SSN format
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Now we can use the compiled pattern multiple times efficiently
	fmt.Println(r.MatchString("123-45-6789")) // true
	fmt.Println(r.MatchString("123-456-789")) // false

	// Find all matches in a text
	text := "SSN: 123-45-6789 and 987-65-4321 and 123-453-678ww9"
	matches := r.FindAllString(text, -1)
	fmt.Println("Found SSNs:", matches)
}
