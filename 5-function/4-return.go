package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x // 可以有多个返回值
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}
