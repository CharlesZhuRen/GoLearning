package main

import (
	"fmt"
)

/*
如果key不存在，会返回该value类型的默认值，但是程序不会报错
所以需要通过使用ok-idiom获取值，判断key或value是否存在
*/

func main() {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"] // x为value（若不存在则为默认值），ok为是否存在
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)
}
