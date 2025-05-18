package main

import (
	"fmt"
)

func main() {

	a := [7]string{"Mon", "Tues", "Weds", "Thurs", "Fri", "Sat", "Sun"}

	// Print days using printf
	fmt.Println("\n----- Printing days using printf -----")
	for index, day := range a {
		fmt.Printf("Day(%d) = %s\n", index+1, day)
	}

	// Print days with position using printf
	fmt.Println("\n----- Printing days with position -----")
	for i := 0; i < len(a); i++ {
		fmt.Printf("Position %d contains %s\n", i, a[i])
	}

	cs := [5]string{"Britian", "Canada", "Australia", "India", "New Zealand"}

	cs1 := cs[1:4]
	cs2 := cs[:3]
	cs3 := cs[2:]
	cs4 := cs[:]

	fmt.Println("------ countries------")
	fmt.Println("array countries =", cs)
	fmt.Println("cs1=", cs1)
	fmt.Println("cs2=", cs2)
	fmt.Println("cs3=", cs3)
	fmt.Println("cs4=", cs4)

	fmt.Println("------- end countries --------")

	slice1 := a[1:]

	slice2 := a[3:]

	fmt.Println("-------Before Modification-----")
	fmt.Println("a=", a)
	fmt.Println("slic1=", slice1)
	fmt.Println("slice2=", slice2)
	slice1[0] = "PINK"
	slice1[1] = "WEDS"
	slice1[2] = "THURS"
	slice2[1] = "FRI"

	fmt.Println("-----after modification----")
	fmt.Println("a=,", a)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)

}
