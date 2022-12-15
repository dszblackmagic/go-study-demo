package main

import "fmt"

// 从前面的案例中可以知道，接口可以同时可以被多个对象实现，当然一个对象也可以同时实现多个接口，如下

// Move 移动
type Move interface {
	Move()
}

// Sound 声音
type Sound interface {
	Sound()
}

type Dog struct {
	name string
}

// 一个 struct 实现两个接口

func (d Dog) Move() {
	fmt.Println(d.name, "正在移动...")
}
func (d Dog) Sound() {
	fmt.Println(d.name, "正在发出声音...")
}

func main() {

	// 定义接口
	var m Move
	var s Sound
	// 创建实例
	dog := Dog{"小黑狗"}
	m = dog
	s = dog
	// 调用方法
	m.Move()
	s.Sound()
}
