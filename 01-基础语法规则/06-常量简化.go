package main

import "fmt"

// 定义常量需要用到的关键字为 const
// 常量定义后不可改变
const (
	a1 = 10
	b1
	c1 = 20
	d1
)

func main() {

	fmt.Println("输出定义的常量:", a1, b1, c1, d1)
}
