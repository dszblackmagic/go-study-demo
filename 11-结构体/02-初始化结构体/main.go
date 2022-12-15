package main

import "fmt"

// 我们除了可以通过 new() 来创建结构体的实例子，我们还可以对结构体进行初始化
// 定义一个结构体

type Person struct {
	name string
	age  int
	sex  string
}

func main() {

	// 初始化结构体方式一
	// 这种方式我们可以直接初始化值,但是初始化的顺序要和结构体中字段的顺序一直，并且初始化的字段数必须与结构体字段数相同
	p1 := Person{"张三", 18, "男"}
	fmt.Println(p1)
	// 初始化方式二：我们可以只初始化部分
	// 这种初始化方式可以自己指定初始化的字段，没初始化的字段会按照类型的默认值来赋值
	p2 := Person{name: "李四", sex: "男"}
	fmt.Println(p2)
}
