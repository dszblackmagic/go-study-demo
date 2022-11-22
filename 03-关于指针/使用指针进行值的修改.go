package main

import "fmt"

func main() {

	// 设定两个变量
	a := 10
	b := 20
	// 调用没使用指针的方法后输出结果
	modifyValue(a)
	fmt.Println("修改后a的值为：", a)
	// 调用使用指针对值进行修改的函数
	modifyFromPoint(&b)
	fmt.Println("修改后b的值为：", b)
	// 从结果可以得知，a的变量值未发生改变，b的变量值发生了改变
}

// 当前函数是未使用指针
func modifyValue(num int) {
	num = 1000
	fmt.Println("modifyValue方法: ", num)
}

// 当前函数传递的是指针
func modifyFromPoint(num *int) {
	// 使用指针地址所指向的修改值
	*num = 100
	fmt.Println("modifyFromPoint方法: ", *num)
}
