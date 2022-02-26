package main

// 格式：for init; condition; post { }

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b { // 这个就跟while一样了 太色了
		a++
		fmt.Printf("a 的值为: %d\n", a)
		if a > 10 {
			break
		}
	}

	for i, x := range numbers { // 这个又跟python有点像 太色了
		if x == 5 {
			continue
		}
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
