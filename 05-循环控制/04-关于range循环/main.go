package main

import "fmt"

// for range 是for循环中的特有循环结构
// 可以遍历数组、切片、字符串、map 及管道（channel），在后续的案例中我们也会经常用到它现在我们可以先试一试
func main() {

	// 案例：循环遍历字符串中的每一个字
	// 定义一个字符串
	str := "凳子学习之旅"
	// for range的语法解析, key实际上就是当前被遍历的内容的索引，value则是值，
	/**
	语法：for key, val := range 要遍历内容 {}
	*/
	for index, char := range str {
		fmt.Printf("当前的位置是：%d , 当前的值是: %c \n", index, char)
	}
}
