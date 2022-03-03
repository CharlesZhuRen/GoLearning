package main

import "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

func main() {
	var a = first{1, false, 2} // 对其所有属性进行初始化
	var b *first = &a
	fmt.Println(a.b, a.a, a.name, &a, b.a, &b, (*b).a, b)

	if b != nil {
		fmt.Println("not nil")
	}
}
