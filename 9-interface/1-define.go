package main

import (
	"fmt"
)

// an interface can be implemented by any instance
// an instance can implement any number of interfaces
// any type has implemented a nil interface (interface which contains 0 method)

type Phone interface { // an interface named Phone, contains a method named call
	call() // if an instance implements all methods in this
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() { // the NokiaPhone struct implements the call method
	// so it implements the interface
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() { // the iPhone struct implements the call method
	// so it implements the interface
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone // declare an interface

	phone = new(NokiaPhone) // instance the interface
	phone.call()

	phone = new(IPhone)
	phone.call()

}
