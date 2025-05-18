package main

import (
	"fmt"
	"sort"
)

func main() {

	//slice1 := make([]string, 3, 10)
	//copy(slice1, []string{"red", "blue", "green"})
	slice1 := []string{"Red", "Orange", "Yellow"}

	slice2 := []string{"Blue", "Green", "Violet"}

	slice3 := append(slice1, slice2...)

	primes := []int{2, 5, 9, 11, 7, 1, 323, 27, 63, 91}
	sort.Ints(primes)

	animals := []string{"cows", "wildbeest", "bear", "panther", "tiger"}

	sorted := sort.StringsAreSorted(animals)

	fmt.Printf("slice1 = %v, len = %d, cap=%d \n", slice1, len(slice1), cap(slice1))

	fmt.Printf("slice2 = %v, len = %d, cap=%d \n", slice2, len(slice2), cap(slice2))

	fmt.Printf("slice3 = %v, len = %d, cap=%d \n", slice3, len(slice3), cap(slice3))

	slice1[0] = "pink"

	fmt.Println("\nslice1=", slice1)

	fmt.Println("\nSlice2=", slice2)

	fmt.Println("\nSlice3=", slice3)

	fmt.Println(primes)

	fmt.Println(sorted)
	fmt.Println(animals)

	sort.Strings(animals)

	sorted = sort.StringsAreSorted(animals)

	fmt.Println(sorted)
	fmt.Println(animals)

	fmt.Println("Before sorting:", animals)

	// Reverse sort (descending order)
	sort.SliceStable(animals, func(i, j int) bool {
		return animals[i] < animals[j]
	})

	fmt.Println("After reverse sorting:", animals)

	jonesFamily := []struct {
		Name string
		Age  int
	}{
		{Name: "Caesaria", Age: 22},
		{Name: "Augustus", Age: 25},
		{Name: "Julius", Age: 3},
	}

	sort.SliceStable(jonesFamily, func(i, j int) bool {
		return jonesFamily[i].Age > jonesFamily[j].Age
	})

	fmt.Println(jonesFamily)

	num := int{36, 45, 52, 29}

	sort.Sort(sort.Reverse(sort.Intslice(num)))

	fmt.Println(num)

}
