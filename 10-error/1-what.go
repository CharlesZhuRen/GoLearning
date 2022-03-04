package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil { // 处理错误的惯用方法
		fmt.Println(err)
		return
	}
	//根据f进行文件的读或写
	fmt.Println(f.Name(), "opened successfully")
}
