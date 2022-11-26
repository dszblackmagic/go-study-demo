package main

import "fmt"

func main() {
	// 创建变量
	str := "this is str"
	// 将str的地址赋予到新的变量
	ptr := &str
	// 已知 ptr 的值为 str 的地址，因此我们要通过 ptr 来获取 str 的值
	// 利用指针取值需要用到 * 号
	fmt.Println("利用指针地址获取的值为：", *ptr)

}
