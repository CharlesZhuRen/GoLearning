package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin) // 创建reader对象
	s1, _ := reader.ReadString('\n')    // 这个_是啥意思？哦哦 是error
	fmt.Println("读到的数据：", s1)

}
