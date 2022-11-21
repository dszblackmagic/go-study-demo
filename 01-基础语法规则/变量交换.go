package main

import "fmt"

func main() {

	// 定义变量a、b，并且分别赋值10、20
	a, b := 10, 20
	// 第一轮输出查看结果
	fmt.Println("原本a、b的结果为：", a, b)
	// 然后我们可以直接用 = 号来对a、b进行快速赋值，交换他们的位置
	a, b = b, a
	fmt.Println("新输出的a、b结果为：", a, b)

}
