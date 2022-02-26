package main

import "fmt"

// 常量组：不用居然也不会报错哎
const (
	Unknown = 0
	y       // 如果不指定类型和初值，则和Unknown相同
	Female  = 1
	Male    = 2
	aha     = "aha"
	s       // 如果不指定类型和初值，则和aha相同
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
}
