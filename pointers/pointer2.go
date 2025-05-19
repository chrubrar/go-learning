package pointers

import "fmt"

func main() {

	var num int = 42
	var ptr = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value at address stored in ptr:", *ptr)

	//*ptr = 100
	var ptr1 *int = &num
	*ptr1 = 100
	fmt.Printf("\nAfter *pointer = 100:\n")
	fmt.Printf("num's new value: %d\n", num)
	fmt.Printf("ptr1's value: %p\n", ptr1)
	fmt.Printf("ptr1's value: %p\n", ptr1)
}
