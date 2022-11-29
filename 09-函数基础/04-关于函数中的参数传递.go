package main

import "fmt"

// 在参数传递的过程中分别为：值传递、引用传递
// 值传递：在函数中改变参数不会影响实际参数
// 引用传递：在函数中改变参数会影响实际参数
func main() {

	// 定义一个string 变量
	str := "这是一个string"
	// 调用函数
	valueFunc(str)
	// 输出结果查看是否实际参数有在函数中被影响到
	fmt.Println("输出str：", str)
	// 定义一个int函数
	x := 10
	addrFunc(&x)
	fmt.Println("得到x的结果为：", x)
	// 定义一个int切片
	nums := make([]int, 10)
	// 调用相关函数
	listFunc(nums)
	// 输出切片结果
	fmt.Println(nums)
}

// 值传递
func valueFunc(str string) {
	// 改变参数
	str = "在valueFunc中改变参数"
}

// 引用传递
func addrFunc(x *int) {
	*x = 100
}

// map、slice、chan、指针、interface默认以引用的方式传递。
// 注意：实际上在使用切片的时候最好采用地址传递，而不是直接传值，因为会在扩容的过程中产生地址变动（比如你修改的值超过来原本make的cap容量）
// 这样会而导致结果出现问题
func listFunc(args []int) {
	// 重新初始化
	for i := 0; i < 5; i++ {
		args[i] = i
	}
}
