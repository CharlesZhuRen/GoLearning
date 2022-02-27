package main

import "fmt"

/*
数组是值类型的
所以在被分配给一个新变量后，对新变量的改动不会影响到原数组

数组的大小是类型的一部分 所以[5]int和[25]int是不同的类型
也因此，数组不能被调整大小
*/

func main() {
	a := [...]string{"Singapore", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "USA"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
