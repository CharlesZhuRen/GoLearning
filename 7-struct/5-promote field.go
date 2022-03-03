package main

import (
	"fmt"
)

type Address1 struct {
	city, state string
}

type Person1 struct {
	name string
	age  int
	Address1
}

func main() {
	var p Person1
	p.name = "Naveen"
	p.age = 50
	p.Address1 = Address1{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
