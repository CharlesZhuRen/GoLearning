package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点类型：")
	fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	// 跟c一摸一样 说实话 感觉有点蠢
	fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

	// 神奇的是，中间有空格有符号居然都可以忽略，所以scan就是只读该读的东西吗
}
