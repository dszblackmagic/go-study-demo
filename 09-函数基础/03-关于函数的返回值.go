package main

import "fmt"

// 在go语言中函数可以有多个返回值，并且可以对返回值进行命名然后在函数内直接使用(当成一个局部变量)，然后用return进行默认返回
func main() {

	// 调用多返回值的函数
	x, y := numFunc(10, 20)
	fmt.Println("得到的结果为:", x, y)
	// 调用带有返回值命名的函数
	fmt.Println("---------分割线----------")
	names, m, num := stringFunc()
	fmt.Println(names, m, num)
}

// 在函数参数后面用(返回值列表..)
func numFunc(x, y int) (int, int) {
	// 将x,y分别扩大十倍后返回
	return x * 10, y * 10
}

// 对函数的返回值进行命名(被命名的返回值可以直接当成一个局部变量来进行使用)
func stringFunc() (names []string, m map[string]int, num int) {
	// 可以直接重新给names变量赋值
	names = []string{"a", "b", "c"}
	// 也可以往map对象中添加内容（需要先赋予空间）
	m = make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	// 或者给num进行赋值
	num = 10
	// 直接用return将这些值直接进行默认返回
	return

}
