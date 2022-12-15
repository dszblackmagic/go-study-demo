package main

import "fmt"

// 空接口实际上是可以作为参数的，从上一个案例中我们知道空接口等于可以被任何类型实现
// 所以当以空接口作为参数的时候，其实就说当前的函数可以接受任何类型的参数
func main() {

	// 可以定义不同的变量进行传参调用
	x := 10
	y := "这是str"
	z := true
	// 调用
	interfaceVar(x)
	interfaceVar(y)
	interfaceVar(z)

	// 当然还可以将接口作为 map 的类型，这样的 map 值可以传递任何类型
	var person = make(map[string]interface{})
	// 因为值是 interface 类型， 所以在赋值的时候可以赋予任何值
	person["name"] = "墨菲特"
	person["age"] = 18
	person["judge"] = true
	// 输出
	fmt.Println(person)
}

func interfaceVar(x interface{}) {
	fmt.Printf("获取到的类型为: %T, 获取到的值为: %v \n", x, x)
}
