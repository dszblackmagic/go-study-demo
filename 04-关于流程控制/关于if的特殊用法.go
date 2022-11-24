package main

import "fmt"

// 前面的案例说的都是if常见的写法也比较好理解
// 接下来说的是if中比较特殊的用法，主要体现在可以在条件判断的时候进行赋值操作
func main() {

	// 定义一个变量
	a := 10
	// 在if中我们去判断 b 变量，但是这个 b 变量是只在 if 条件语句中存在
	if b := a; b < 100 {
		// 注意在赋值与条件判断之间要用分号隔开
		fmt.Println("我们得到结果,你赋予b的值小于100")
	}
}
