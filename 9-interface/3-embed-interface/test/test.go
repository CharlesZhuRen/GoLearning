package test

import (
	"fmt"
)

type Controller struct { // define a struct
	M int32
}

type Something interface { // define an interface and its methods
	Get()
	Post()
}

func (c *Controller) Get() { // use the controller struct to implement the get method in the interface
	fmt.Print("GET")
}

func (c *Controller) Post() { // implement the other method
	fmt.Print("POST")
}

// then all methods in the interface has been implemented by the struct
// so that we can say that the struct implements the interface
