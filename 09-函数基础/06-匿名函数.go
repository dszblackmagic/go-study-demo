package main

import "fmt"

// 在go中我们可以将函数当作变量来使用
func main() {

	// 定义一个变量，返回由匿名函数来决定
	str := func() string {
		return "这是匿名函数的返回值"
	}
	// 还有一种情况我们可以在匿名函数定义的时候就默认进行调用，如下
	// 尾巴的() 里面可以进行传参
	func(data int) {
		fmt.Println("不需要返回值直接调用匿名函数:", data)
	}(100)
	// 获取结果,这时候的str 变量要当成一个函数来使用 所以是 str()
	fmt.Println(str())

	fmt.Println("------分割线-------")
	// 可以将匿名函数作为回调函数
	visit([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println("得到结果:", v)
	})

}

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
