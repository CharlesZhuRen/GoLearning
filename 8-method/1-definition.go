// why it is different from func:
// it has a receiver
// this is the only different part

package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

// 这不就相当于是类的方法吗》
// 只不过这里的类被称为receiver而已

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
}
