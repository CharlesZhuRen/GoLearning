package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	for _, v := range a { //如果只需要值并希望忽略索引，就用_（称作blank标识符）替换索引
		fmt.Printf("%.2f\n", v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
