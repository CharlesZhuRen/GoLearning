package main

import (
	"fmt"
)

// 空白标识符：_ 可以代替任何类型的任何值
// 其实在我的python编程里就经常用的

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
	// 虽然不会用到，但还是得接收，于是就用空白标识符接一下
	fmt.Printf("Area %f ", area)
}
