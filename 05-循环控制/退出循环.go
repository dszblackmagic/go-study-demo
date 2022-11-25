package main

import "fmt"

// 我们可以使用一些关键字来中途退出循环
func main() {

	// 假设我们需要在当循环运行到 i = 5 的时候退出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			// 当i=5的时候就退出当前循环不再继续执行下去了
			// 需要注意的是，一个break只能退出一层循环
			// 也就是说如果你的for是双层嵌套循环，你需要在外层的for中也进行条件判断才能退出外层的循环
			break
		}
		fmt.Println("break退出前输出当前的值为:", i)
	}

	// 或者我们可以使用 continue来跳过当前循环
	// 假设我们需要跳过1-10之中的4，输出其他的内容我们就可以这样操作
	for i := 0; i < 10; i++ {
		if i == 4 {
			// 如果 i=4 则跳过不输出
			continue
		}
		fmt.Println("continue未跳过的内容为:", i)
	}
}
