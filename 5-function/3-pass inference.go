package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
// 只有知道a所在的地址，才能修改它的值，所以这里要传指针进去

// tips: 传指针比较轻量级，所以如果要传的东西是个比较大的结构体，那么直接传指针会更高效一点
// slice, map等的实现机制类似于指针，所以可以直接传递，而不需要取地址之后传指针
// 但如果涉及到修改slice的长度，那么仍然需要传指针

func add1(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}
func main() {
	x := 3
	fmt.Println("x = ", x)    // 应该输出 "x = 3"
	x1 := add1(&x)            // 调用 add1(&x) 传x的地址
	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
}
