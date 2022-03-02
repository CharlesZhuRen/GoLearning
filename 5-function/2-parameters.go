package main

// 支持可变参数：func myfunc(arg ...int) {}
// 这里的arg是一个int的slice
//for _, n := range arg {
//fmt.Printf("And the number is: %d\n", n)
//}

import (
	"fmt"
	"math"
)

func main() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9)) // pass a value as parameter

}
