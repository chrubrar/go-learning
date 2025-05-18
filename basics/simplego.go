package basics

import (
	"fmt"
	
)


func Max(num1, num2 int) int {
	var result int
	
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}
