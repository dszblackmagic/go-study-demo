package main

import "fmt"

// 在学习过数组的定义和取值后，我们可以学会遍历数组来将内容进行输出，而不是一个一个的取值
func main() {

	// 定义数组
	arr := [...]int{1, 2, 3, 4}
	// 遍历数组的方式一
	// len(arr)来获取数组的长度来决定循环的退出条件
	for i := 0; i < len(arr); i++ {
		fmt.Println("循环一输出数组值:", arr[i])
	}
	fmt.Println("--------分割线---------")
	// 第二种循环遍历数组
	for index, value := range arr {
		fmt.Println("循环二数组的数组索引：", index, " 值：", value)
	}
}
