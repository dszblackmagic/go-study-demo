package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 定义一个字符串
	str := "聪明的墨菲特"
	// 获取到内容c 并且进行输出
	fmt.Println(str[1])
	// 输出的是字节的长度
	fmt.Println(len(str))
	// 想要获取utf8的字符长度则需要使用utf8支持的方法
	fmt.Println("得到字符串长度为:", utf8.RuneCountInString(str))

}
