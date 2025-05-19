package datastructures

import "fmt"

func main() {

	type Compensation struct {
		baseSalary      int
		healthInsurance int
		bonus           int
	}

	type Employee struct {
		firstName    string
		lastName     string
		age          int
		fullTime     bool
		compensation Compensation
	}

	ben := Employee{firstName: "Ben", lastName: "Dysinger", age: 25, fullTime: true, compensation: Compensation{baseSalary: 50000, healthInsurance: 500, bonus: 1000}}
	fmt.Println("ben =", ben)
	//ben.compensation.bonus = 2100

	fmt.Println("ben with bonus =", ben.compensation.bonus)
}
