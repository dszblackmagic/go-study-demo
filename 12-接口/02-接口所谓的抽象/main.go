package main

import "fmt"

// 我们知道我们可以把接口想成是某种类型对象共有的抽象行为，比如每个人都会：吃喝玩乐，
// 但是每个人又会有自己的特长比如：游泳、拳击、编程等等
// 我们可以把有共性的事情抽取出来进行实现，interface 是一组 method 的集合，接口做的事情就像是定义一个协议（规则）
// 以下案例：不同职业运动员有各自的特长，有拳击运动员特长为：打拳，游泳运动员特长为：游泳。用代码的形式表现如下

type Athlete interface {
	//BeGoodAt 擅长
	BeGoodAt(name string) error
}

type Boxer struct {
	name string
}

type Swimmer struct {
	name string
}

// 分别实现对应的方法

func (b *Boxer) BeGoodAt(name string) error {
	// 注意：接口中的方法必须全部被实现，这个接口才可以正常被编译
	fmt.Println(name, "擅长打拳击！")
	return nil
}

func (s *Swimmer) BeGoodAt(name string) error {
	fmt.Println(name, "擅长游泳！")
	return nil
}

func main() {

	// 定义两个实例
	boxer := Boxer{"李白"}
	swimmer := Swimmer{"杜甫"}
	// 创建接口
	var ath Athlete
	// 将接口进行绑定调用查看结果
	ath = &boxer
	ath.BeGoodAt(boxer.name)
	ath = &swimmer
	ath.BeGoodAt(swimmer.name)

}
