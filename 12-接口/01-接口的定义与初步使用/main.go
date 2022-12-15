package main

import "fmt"

// 接口：可以理解为接口提供了一个方式来说明对象的行为，也可以理解成将一些对象的行为共性抽取出来
/*
举个例子：我们知道每个动物都可以发出叫声，但是每个动物具体的叫声不同。为此我们可以将叫声这个行为抽出来，
在遇到对应的对象的时候再具体实现它的叫声即可。这样的方式可以更加的灵活，同时 Go 语言也希望能面向接口编程。
*/

/*
定义接口：关键字 - interface
*/

// Animal 我们定义一个叫动物的方法
type Animal interface {
	// 定义方式的格式：方法名(参数列表) 返回值类型

	// Cry 在接口里面来定义个行为是：叫
	Cry() string
}

// Cat 我们可以定义个结构体为猫
type Cat struct {
	name string
}

// 实现接口

func (c *Cat) Cry() string {

	return c.name + "发出叫声：喵喵喵～"
}

func main() {

	// 首先我们创建一个猫的实例
	cat := new(Cat)
	// 我们再创建一个接口变量
	var animal Animal
	// 将猫这个实例赋予接口（因为在猫这个struct的接收器中已经有去实现 Cry 这个方法了，我们需要将它与接口绑定上）
	animal = cat
	// 这时候我们就可以调用接口中的 Cry() 方法
	fmt.Println(animal.Cry())
}
