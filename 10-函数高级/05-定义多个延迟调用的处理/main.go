package main

import (
	"fmt"
	"log"
)

// 我们上一个案例说了延迟调用是在函数结束的最后进行注册使用的，如果我们同时定义多个延迟调用会怎么样？
func main() {

	// 定义普通函数
	fmt.Println("多延迟调用的开始...")
	// 假设我们定义三个延迟调用
	// 因为延迟调用实际上是逆排序的，所以就像一个栈一样，第一个进去的最后出来，最后一个定义的则第一个出来
	defer log.Println("延迟调用一")
	defer log.Println("延迟调用二")
	defer log.Println("延迟调用三")

	// 结束输出
	fmt.Println("多延迟调用的结束....")
}
