package main

import "fmt"

// 函数：理解成一个可以重复使用的代码片段
/**
函数的定义方式：
func function_name( [参数列表] ) [返回值类型] {
   函数体
}
注意： 返回值和参数都不是必要的，在不同的使用场景下是有可能不需要的
*/
func main() {

	// 写一个函数，可以返回两整数相加的结果
	// 定义两个int变量
	x := 10
	y := 20
	// 调用函数： 函数名(需要的参数)
	// res 为接受函数返回值的变量
	res := sum(x, y)
	// 输出结果
	fmt.Println("调用函数得到的结果为:", res)

}
func sum(x, y int) int {
	// 返回值则是x与y的和
	return x + y
}
