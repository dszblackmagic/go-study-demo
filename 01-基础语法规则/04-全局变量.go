package main

import "fmt"

// 定义全局变量，必须有var关键字，并且在函数体外面
var a = 10
var b = "let's go"

func main() {
	// 在main()函数中调用a、b
	fmt.Println("输出全局变量为：", a, b)
	sum()
}

func sum() {
	// 因为是全局使用，在要在这个源码中我们的方法都能调用到
	// 因此我在sum()函数中也可以调用到b变量的值
	fmt.Println("调用sum()函数输出的全局变量为:", b)
}
