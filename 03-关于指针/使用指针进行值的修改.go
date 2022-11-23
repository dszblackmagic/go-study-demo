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

	x, y := 1, 2
	// 交换指针地址
	swapPointAddr(&x, &y)
	fmt.Println("交换指针地址后的值为：", x, y)
	// 形参修改指针地址, 原参数的指针地址不变。
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

func swapPointAddr(num, num2 *int) {
	fmt.Println("swapPointAddr交换指针地址为：", num, num2)
	// 将a,b对应的指针地址互换
	num, num2 = num2, num
	fmt.Println("swapPointAddr交换指针地址后的值为：", *num, *num2)
}
