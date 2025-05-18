package greeting

import (
	"fmt"
)

var Pet = "Powderpuff"

type Unicorn struct {
	Name  string
	Color string
}

func (u Unicorn) String() string {
	return fmt.Sprintf("The Unicorn's name is %s and its color is %s", u.Name, u.Color)

}

func Hello() {
	fmt.Println("Hello, World!")
}
