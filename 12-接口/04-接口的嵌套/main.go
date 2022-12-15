package main

import "fmt"

// 我们除了单独定义接口外，接口与接口之间是可以嵌套使用的。嵌套创造出来的新接口使用方式仍然一样，就是要实现其中的内容即可。

type Mover interface {
	Mov()
}
type Sayer interface {
	Say()
}

// 定义一个接口同时包含以上两个接口的内容

type Nesting interface {
	Mover
	Sayer
}

// 想要使用 Nesting 这个接口，同样的只要实现里面的接口的方法即可

type Nest struct {
	name string
}

func (n Nest) Say() {
	fmt.Println(n.name, "正在Say...")
}

func (n Nest) Mov() {
	fmt.Println(n.name, "正在移动....")
}

func main() {

	// 定义实例
	n := Nest{"凳子"}
	// 定义接口
	var ni Nesting
	// 调用
	ni = n
	ni.Say()
	ni.Mov()

}
