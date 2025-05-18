package main

import "fmt"

type Employee struct {
	lastName, firstName string
	salary              int
	age                 int
	fullTime            bool
	vacationDays        int
}

func main() {

	buffy := Employee{firstName: "Buffy", lastName: "VampireSlayer", salary: 50000, age: 25, fullTime: true, vacationDays: 20}

	fmt.Println("buffy =", buffy)
	buffy.vacationDays = 22
	fmt.Println("buffy =", buffy)

	giles := struct {
		lastName, firstName string
		salary              int
		age                 int
		fullTime            bool
		vacationDays        int
	}{firstName: "Giles", lastName: "Angel", salary: 50000, age: 25, fullTime: true, vacationDays: 20}

	fmt.Println("giles =", giles)
}
