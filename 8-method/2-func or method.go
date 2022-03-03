package main

import (
	"fmt"
)

// go is not a pure OOP language, it doesn't support class
// so, method is kind of implementation of class's behavior

// functions with same func name can be defined based on different types

type Employee1 struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method converted to function with Employee as parameter
*/
func displaySalary(e Employee1) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee1{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp1)
}
