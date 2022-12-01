package main

import "fmt"

// 利用闭包完成一个累加器的案例
func main() {

	// 创建一个计数器
	count := addCount(0)
	// 调用计数器函数
	// 对引用的 addCount 参数变量进行累加，注意 count 不是第 21 行匿名函数定义的，但是被这个匿名函数引用，所以形成闭包
	fmt.Println(count())
	// 不断的调用
	fmt.Println(count())
	// 创建计数器2
	// 对引用的 addCount 参数变量进行累加，注意 count 不是第 21 行匿名函数定义的，但是被这个匿名函数引用，所以形成闭包
	count2 := addCount(10)
	// 调用函数,输出结果
	fmt.Println(count2())
}

// 定义一个函数，每次调用该函数的时候，内部调用一次闭包, 返回值由闭包生成
func addCount(count int) func() int {
	// 在内部返回值是一个闭包
	return func() int {
		count++
		return count
	}
}
