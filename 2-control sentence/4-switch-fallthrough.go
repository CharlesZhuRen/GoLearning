package main

import (
	"fmt"
)

type data [2]int

//var data[2]int  // 这样居然也行。。给我整不会了

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough // 强制不break，即便满足也不跳出switch，继续往下走 但经过测试 只会往下走一个
	case 6:
		x += 20
		fmt.Println(x)
		fallthrough
	case 10086:
		x += 666
		fmt.Println(x)
	}

}
