package main

import "fmt"

type Books2 struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books2 /* 声明 Book1 为 Books 类型 */
	var Book2 Books2 /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "charleszhuren"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "zhuren"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook2(Book1)

	/* 打印 Book2 信息 */
	printBook2(Book2)
}

func printBook2(book Books2) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
