package main

import "fmt"

type FullNameType func(string, string) string

type EmployeeFunc struct {
	firstName string
	lastName  string
	age       int
	fullTime  bool
	FullName  FullNameType
}

func main() {

	e := EmployeeFunc{
		firstName: "Ben",
		lastName:  "Dysinger",
		age:       25,
		fullTime:  true,
		FullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}

	fmt.Println("e =", e)
	fmt.Println("e fullName =", e.FullName(e.firstName, e.lastName))
	fmt.Println("e.FullName =", e.FullName("Ben", "Smith"))
}
