package main

import (
	"fmt"
	"github.com/chrubrar/go-learning/greeting"
)

func main() {
	greeting.Hello()
	fmt.Println(greeting.Pet)

	uni := greeting.Unicorn{
		Name:  "Sparkle",
		Color: "Rainbow",
	}
	fmt.Println(uni)
}
