package main

import "fmt"

// 可以利用闭包来完成一个生成器
func main() {

	// 通过生成器创建一个游戏角色
	generator := game("凳子先生")
	name, hp := generator()
	// 输出结果
	fmt.Println("创建角色为:", name, "｜初始血量为:", hp)
}

// 创建一个生成器的函数,返回闭包
func game(name string) func() (string, int) {

	// 创建一个血量值
	hp := 150
	// 创建一个闭包函数
	return func() (string, int) {
		return name, hp
	}
}
