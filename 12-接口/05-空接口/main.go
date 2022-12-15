package main

import "fmt"

// 什么是空接口？ 空接口指的是没有定义任何方法的接口。
// 因为空的接口没有任何方法，所以就意味着可以理解成任何类型都实现了这个接口。
func main() {

	// 定义一个空接口
	var x interface{}
	// 因为任何类型都可以实现空接口，所以可以赋予任何类
	str := "let's go"
	x = str
	// 输出一下 x 的类型和值看看
	fmt.Printf("当前类型为 %T, 当前值为 %v", x, x)

}
