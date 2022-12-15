package main

import "fmt"

// 我们原本定义结构体的时候，结构体是有结构体名称的，其实和函数一样，结构体也可以是匿名的
func main() {

	// 定义一个匿名结构体变量如下
	structDemo := struct {
		// 这里是匿名结构体的字段
		name string
		age  int
	}{
		// 可以在这个括号内初始化字段值
		// 方式一
		//name: string("凳子"),
		//age:  int(20),

		// 方式二
		"凳子",
		10,
	}

	// 直接调用这个匿名结构体
	fmt.Println("名字为：", structDemo.name, "年龄为：", structDemo.age)

}
