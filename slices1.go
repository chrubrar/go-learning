package main

import (
	"fmt"
)

func main() {

	days := []string{"Thurs", "Fri", "Sat", "Sun"}

	for index, day := range days {
		fmt.Printf("Day(%d) = %s\n", index+1, day)
	}

	for i := 0; i < len(days); i++ {
		fmt.Printf("Position %d contains %s\n", i+1, days[i])
	}

	primenumbers := []int{1, 2, 3, 5, 7, 11, 13, 17, 19}

	for index, number := range primenumbers {
		fmt.Printf("Primennumber(%d) = %d\n", index+1, number)
	}
}
