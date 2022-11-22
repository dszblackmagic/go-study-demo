package main

import "fmt"

func main() {
	// 指针的取值语法为 在变量名前面添加&操作符（前缀）来获取变量的内存地址
	// 定义一个变量
	dog := "this is dog"
	// 取地址进行输出
	fmt.Printf("将取到的地址输出: %p", &dog)
}
