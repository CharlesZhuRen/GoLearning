package main

import "GoLearning/7-struct/8-export/computer"
import "fmt"

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}
