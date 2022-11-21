package main

import (
	"fmt"
	"strings"
)

// 关于字符串进行位置的索引与查找
// strings.Index()： 正向搜索子字符串
// strings.LastIndex()：反向搜索子字符串
func main() {

	str := "桌子椅子凳子,go语言学习"
	index := strings.Index(str, ",")
	fmt.Println("逗号前定位:", index)
	// index: 表示的是输出从index当前这个位置的内容到最后（从头到尾）
	fmt.Println(str[index:])
}
