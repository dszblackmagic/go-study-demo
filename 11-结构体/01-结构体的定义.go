package main

import "fmt"

// Go语言可以通过自定义的方式形成新的类型，结构体可以理解成就是一个复合类型（零到多个元素且不同类型的任意值组合而成）
/*
Go语言中的结构体语法如下：
type identifier struct {
    field1 type1
    field2 type2
    ...
}
*/
// 根据结构体的语法我们可以创建一个结构体

type People struct {
	// 注意结构体中的字段如果字段名大写则可以被其他包调用，如果小写则只能在本包中调用
	name string
	age  int
}

func main() {

	// 我们可以通过 new() 的方式来创建这个结构体（将这个结构体变成一个实例子）
	p := new(People)
	// 给结构体中的字段赋值
	p.name = "凳子"
	p.age = 18
	// 输出结果
	fmt.Println("人的名称为:", p.name, " 人的年龄为：", p.age)

	// 第二种创建结构体的方式可以采用取地址的方式，对等于 new() 的方式
	pp := &People{}
	pp.name = "测试另一种结构体定义方式"
	pp.age = 20
	fmt.Println("人的名称为:", pp.name, " 人的年龄为：", pp.age)

}
