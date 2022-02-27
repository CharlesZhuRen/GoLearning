package main

import (
	"fmt"
)

func main() {
	// if statement, condition
	// 先执行statement，然后执行condition
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
