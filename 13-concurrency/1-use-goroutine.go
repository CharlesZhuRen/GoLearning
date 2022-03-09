package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second) // 如果不加这一行，那么可能不会输出hello world goroutine
	// 不sleep的话，就有可能在main goroutine终止之前执行
	fmt.Println("main function")
}
