package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human // 匿名字段，那么默认Student就包含了Human的所有字段
	// 这里是不是可以理解为继承呢 -> 还真是
	speciality string
}

func main() {
	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息 -> 这里怎么不脱裤子放屁用指针了？
	// 讲道理 其实按照java的思想的话 这里确实得另写一个函数 然后传指针作为参数
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}
