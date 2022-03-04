package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) Show() {
	fmt.Println("Person-->", p.name)
}

//类型别名
type People = Person

type Student2 struct {
	// 嵌入两个结构
	Person
	People
}

func (p People) Show2() {
	fmt.Println("People------>", p.name)
}

func main() {
	//
	var s Student2

	//s.name = "王二狗" //ambiguous selector s.name
	s.People.name = "李小花"
	s.Person.name = "王二狗"
	//s.Show() //ambiguous selector s.Show
	s.Person.Show()
	s.People.Show2()
	fmt.Printf("%T,%T\n", s.Person, s.People) //main.Person,main.Person

}
