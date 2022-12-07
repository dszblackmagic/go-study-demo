package main

import "fmt"

// 在实战的过程中我们通常会将结构体对外不公开，更希望用户能够用工厂的方式来创建结构体（可以方便我们自身在初始化的时候一些自定义操作）
// 注意：小写结构体名即表示不对外公开

type dog struct {
	Name string
	Age  int
}

//定义一个方法,专门用来初始化这个结构体

func NewDog(name string, age int) *dog {
	// 在这里面进行初始化
	// 这样的操作可以防止别的包直接调用结构体，而只能强制使用此方法来进行结构体的初始化
	return &dog{
		Name: name,
		Age:  age,
	}
}

func main() {

	// 初始化对象
	dog := NewDog("田园犬", 10)
	// 输出结果
	fmt.Printf("%+v", dog)
}
