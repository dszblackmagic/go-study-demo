package main

import (
	"fmt"
	"strconv"
)

// 本案例讲述Go中字符串与各个类型之间的常见转换
// 需要用到strconv包
func main() {

	// str 转 int
	str := "10"
	// 对于err 可以选择不接受
	intRes, _ := strconv.Atoi(str)
	fmt.Printf("整型转换得到的结果类型为: %T \n", intRes)

	// string 转 float（浮点）
	strFloat := "3.1415926"
	// 转换 注意ParseFloat该函数的返回值总是float64，如果需要精度为32可以再一次转换
	floatRes, _ := strconv.ParseFloat(strFloat, 32)
	fmt.Printf("浮点转换得到的结果类型为: %T \n", floatRes)

	// int 转 string
	num := 10
	numStr := strconv.Itoa(num)
	fmt.Printf("字符转换得到的结果类型为: %T \n", numStr)

}
