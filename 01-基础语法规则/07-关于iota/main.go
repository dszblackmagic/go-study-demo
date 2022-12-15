package main

import "fmt"

// 定义周一到周天
// iota可以自动将值进行逐步的递增
const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday //6
)

func main() {

	fmt.Println("输出日期的值为：", Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
