package main

import "fmt"

type Bookss struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Bookss /* Declare Book1 of type Book */
	var Book2 Bookss /* Declare Book2 of type Book */

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
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}

// 我的评价是：纯纯的脱裤子放屁
func printBook(book *Bookss) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
