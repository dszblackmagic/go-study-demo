package main

import "fmt"

// 上一个案例中我们讲述了切片在用被切对象的情况下如何进行切片操作
// 我们这个案例中讲述一下如何创建一个空切片
func main() {

	// 创建一个空切片 语法：var 切片名 [] 切片的数据类型
	// 假设我们创建的是一个空的数组切片（任何类型都可以有切片）
	var strSlice []string
	// 判断切片的大小
	fmt.Println("切片的长度为：", len(strSlice))
	// 可以对空切片进行逻辑判断 nil
	fmt.Println("当前切片是否为空:", strSlice == nil)

	// 输出切片
	fmt.Println("输出当前切片:", strSlice)
	// 望切片中添加内容
	// 添加一个元素
	strSlice = append(strSlice, "学习Go语言")
	// 添加多个元素
	strSlice = append(strSlice, "学习如何往切片中添加内容", "再次添加一个元素")
	// 输出添加了元素的切片
	fmt.Println("输出添加元素后的切片：", strSlice)

}
