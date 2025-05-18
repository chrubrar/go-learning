package main

import (
	"fmt"
	"reflect"
)

func printStructFields(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("%s fields:\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
	fmt.Println()
}

func main() {
	type Person struct {
		Name    string
		Age     int
		Address string
		Salary  int
	}

	type Rectangle struct {
		Width  int
		Height int
	}

	rect := Rectangle{Width: 10, Height: 20}
	printStructFields(rect)
	fmt.Printf("Area: %d\n", rect.Width*rect.Height)
	fmt.Printf("Perimeter: %d\n\n", 2*(rect.Width+rect.Height))

	person1 := Person{Name: "John", Age: 30, Address: "123 Main St", Salary: 50000}
	printStructFields(person1)
}
