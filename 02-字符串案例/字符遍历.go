package main

import "fmt"

func main() {

	str := "中文字符"
	// 中文字符只能用for range遍历
	for _, s := range str {
		fmt.Printf("得到的当前中文字符为:%c \n", s)
	}
}
