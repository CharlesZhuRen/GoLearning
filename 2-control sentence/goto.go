package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP // 让我想起了计算机组成原理 这也太色了
			// 这个用法就是用于错误处理
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
