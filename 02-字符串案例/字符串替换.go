package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 桌子椅子凳子,go语言学习 -利用索引将字符串中的go 替换成Java
func main() {

	str := "桌子椅子凳子,go语言学习"
	// 找到'go'
	index := strings.Index(str, "go")
	// 在'go'之前的部分
	start := str[:index]
	// 在'go'之后的部分 (偏移量应该为'go'的长度)
	end := str[index+len("go"):]
	// 将相关的部分进行写入
	var result bytes.Buffer
	result.WriteString(start)
	result.WriteString("Java")
	result.WriteString(end)
	// 输出最终的结果
	fmt.Println("最终得到的结果为: ", result.String())
}
