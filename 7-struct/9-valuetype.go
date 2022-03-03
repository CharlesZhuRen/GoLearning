package main

import (
	"fmt"
)

// if each field in the struct can be compared, then the struct can be compared
// if corresponding fields of 2 structs are equal, then the 2 structs are equal

type name struct {
	firstName string
	lastName  string
}

func main() {
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}
