package main

import (
	"GoLearning/9-interface/3-embed-interface/test"
	"fmt"
)

type T struct {
	test.Controller // the only field in the struct
}

func (t *T) Get() { // rewrite the get method
	//new(test.Controller).Get()
	fmt.Print("T")
}

//func (t *T) Post() {  // rewrite the post method
//	fmt.Print("T")
//}

func main() {
	var something test.Something
	something = new(T)
	var t T
	t.M = 1
	//  t.Controller.M = 1
	something.Get()
	something.Post() // if T doesn't rewrite, then call the original method in Controller
}
