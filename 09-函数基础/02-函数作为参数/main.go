package main

import "fmt"

// 我们可以将作为参数的函数定义为一个新的类型
type pFunc func(s1, s2, s3 string) string

func main() {

	// 调用函数 参数一 传递函数paramFunc, 符合pFunc类型
	s := testFunc(paramFunc, "参数二", "参数三")
	fmt.Println("得到结果为：", s)
}

// 写一个函数,其中第一个参数是一个函数
func testFunc(fn pFunc, s1, s2 string) string {
	return fn("调用参数函数", s1, s2)
}

// 写一个符合传参函数的函数
func paramFunc(s1, s2, s3 string) string {
	// 拼接三个参数的内容
	return fmt.Sprintf("%s %s %s", s1, s2, s3)
}
