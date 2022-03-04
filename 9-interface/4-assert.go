package main

// 当一个函数的形参是interface{}时，在函数中需要对形参进行断言，从而得到它的真实类型

// 安全类型断言
//<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 )

//非安全类型断言
//<目标类型的值> := <表达式>.( 目标类型 )

import "fmt"

func main() {

	var i1 interface{} = new(Student1)
	s := i1.(Student1) //不安全，如果断言失败，会直接panic
	// 这里就因为断言失败panic了

	fmt.Println(s)

	var i2 interface{} = new(Student1)
	s, ok := i2.(Student1) //安全，断言失败，也不会panic，只是ok的值为false
	if ok {
		fmt.Println(s)
	}
}

type Student1 struct {
}
