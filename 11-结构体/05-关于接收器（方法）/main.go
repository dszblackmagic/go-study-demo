package main

import "fmt"

// 其实在 Go 语言中也是支持面向对象的特征，但是它不是一个 OOP 的语言
// 我们定义的结构体，可以拥有它自己特有的方法(比如说我们定义了一个结构体[人]，并且人会[吃饭]，那么这个专属的函数我们可以称之为方法)
// 接收器[方法]的语法：func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {code...}
// 注意：可以发现比起我们普通的函数，多了前面一个()用来定义要绑定的对象(结构体)，所以称之为 方法（接收器）

// 定义一个结构体 People

type Peoples struct {
	name string
	age  int
	sex  string
}

// 定义一个接收器[方法]，来成为 People 的方法
func (p Peoples) eat() {
	// 写一个方法表示, 人正在吃饭
	// 这个方法只有 People 这个类型可以调用
	fmt.Println("正在吃饭...")
	// 并且接收器可以调用对应接收器类型的字段
	fmt.Println("当前的人岁数为：", p.age)
}

func main() {

	// 实例化结构体
	p := new(Peoples)
	// 年龄
	p.age = 20
	// 调用专属的 eat() 方法
	p.eat()

}
