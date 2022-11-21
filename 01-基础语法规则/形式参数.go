package main

import "fmt"

// 定义了全局变量
var aGlobal = "这是全局变量"

func main() {
	// 这里是局部变量
	a := "这是main的局部变量a"
	b := "这是main的局部变量b"
	fmt.Println("输出变量的值为: ", a, b)
	// 在调用test()函数的时候传递局部变量a、b的值
	test(a, b)
}

// 在函数中定义形参a、b，那么当test()函数被调用的时候
// 形参就可以直接被当作局部变量来使用，值则是由调用者的传递内容来决定的
func test(a, b string) {
	// 输出结果
	fmt.Println("这是由main传过来的内容:", a, b)
}
