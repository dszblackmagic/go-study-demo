package main

import "fmt"

// 我们知道在函数中可以传递多个参数，除了一个一个来写，我们如果传递的类型都是相同的话我们可以一次性解决
func main() {

	// 调用函数，传递多个参数
	listNum(1, 2, 3, 4, 5)
	// 调用函数二
	listStringNum("调用函数二", 10, 20, 30, 40, 50)
}

// 利用 ... 就可以表示是多个参数，其本质就是一个slice 切片
func listNum(args ...int) {

	// 遍历参数输出结果
	for i := 0; i < len(args); i++ {
		fmt.Println("listNum结果：", args[i])
	}
}

// 还可以单独穿别的类型参数，后面跟多参数写法，注意 ... 必须在末尾
func listStringNum(str string, args ...int) {
	//输出结果
	fmt.Println("输出多参数结果：", str, args)
}
