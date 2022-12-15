package main

import "fmt"

// Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除
// 即简单来说就是：函数 + 引用环境 = 闭包
func main() {

	// 定义一个变量
	str := "hello go"
	// 定义一个匿名函数去调用str这个变量，注意这时候匿名函数中并没有定义str这个变量
	foo := func() {
		str = "hello Closure"
	}
	// 调用闭包
	foo()
	// 输出结果，str在匿名函数中的改变直接影响到了实际变量
	fmt.Println(str)
}
