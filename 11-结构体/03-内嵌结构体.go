package main

import "fmt"

// 我们知道在结构体中可以定义任何类型的字段，那么意味着我们也可以在结构体中定义其他的结构体作为字段
// 那么我们在调用的时候，如果遇到不同结构体中的重名，需要指定清楚字段名才可以准确的获取到值

type A struct {
	x int
	y int
}

type B struct {
	x int
}

type Demo struct {
	A
	B
}

func main() {
	// 从上面的案例中我们可以看到, 在结构体Demo中我们同时定义了 A、B 两个结构体对象，并且其中字段 x 是重名的
	// 因此我们想要准确的获取到 x 的值， 我们配合字段名准确获取
	// 给 B 中的 x 赋值
	demo := new(Demo)
	demo.B.x = 10
	// 因为 y 是没有重名的，所以是可以直接调用的
	demo.y = 20
	// 输出结果看看
	fmt.Println(demo)

}
