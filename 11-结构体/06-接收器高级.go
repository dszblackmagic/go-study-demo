package main

import "fmt"

// 上一个案例我们知道了接收器，思考一个问题如果接收器类型是 "指针类型" 那么和 "非指针类型" 有什么区别呢
// 主要问题还是和地址有关系，指针是指向地址的，所以在指针类型的接收器中修改内容是会对实际参数产生改变的，查看如下案例

// 定义一直猫

type Cat struct {
	name string
}

// 定义两种类型的接收器，"指针"，"非指针"
func (c Cat) updateName() {
	// 非指针类型的接收器修改名称
	c.name = "花猫"
	fmt.Println("非指针类型修改完成！")
}

func (c *Cat) updateNameByPoint() {
	// 指针类型接收器修改名称
	c.name = "白猫"
	fmt.Println("指针类型修改完成！")
}

func main() {

	// 定义一个结构体实例
	cat := new(Cat)
	// 设置猫的名称
	cat.name = "猫咪"
	// 第一次输出
	fmt.Println(cat.name)
	// 调用非指针类型修改
	cat.updateName()
	fmt.Println(cat.name)
	fmt.Println("----------------")
	// 第二次输出
	fmt.Println(cat.name)
	// 调用指针类型接收器修改
	cat.updateNameByPoint()
	fmt.Println(cat.name)
	// 通过结构可以看到，指针类型的修改是会对实例产生影响的，非指针并不会
}
