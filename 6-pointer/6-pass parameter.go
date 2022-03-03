package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}
func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b) // 接收b 然后直接把b所存放的a的地址上那玩意改成55 太暴力了！！上门查水表属于是
	fmt.Println("value of a after function call is", a)
}
