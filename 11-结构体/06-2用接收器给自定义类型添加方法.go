package main

import "fmt"

// 有了接收器的存在，那么针对于自定义类型，我们就可以给他们添加特定的方法,如下

type List []int

func (l List) getList() {
	for i := 0; i < len(l); i++ {
		// 输出每一个值
		fmt.Println("index:", i, " value:", l[i])
	}
}

func main() {
	// 定义自定义类型（本质是 []int）
	var l List
	l = append(l, 1)
	l = append(l, 2)
	l = append(l, 3)
	// 可以调用特有方法
	l.getList()
}
