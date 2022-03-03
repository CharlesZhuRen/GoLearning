package main

import (
	"fmt"
)

func main() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++ // 一看到这个我就想到计组 一想到计组我就想到绩点
	fmt.Println("new value of b is", b)
	b += 1
	fmt.Println("new value fo b is", b)
}
