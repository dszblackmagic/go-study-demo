package main

import "fmt"

// 在上述的案例中，我们知道了可以通过空接口来传递任何类型的参数。
// 那么有一个问题就是因为传递过来的值可能是任何类型的，我们要如何获取到自己想要的类型？
// 这时候我们就可以通过'断言'的方式来判断传递过来任意类型的数据类型
func main() {

	// 断言语法： 接口变量.(要判断的类型)
	var x interface{}
	// 赋值
	x = "学习go语言"
	// 用断言判断会返回两个值, 第一个值为：当前变量的值 第二个值为：是否想要判断的类型（true/false）
	// 假设我们判断 x 是否为 string类型
	v, judge := x.(string)
	// 输出
	if judge {
		fmt.Println("当前为 string 类型，值为：", v)
	} else {
		fmt.Println("当前类型并不是 string 类型！")
	}
	// 可以用函数来判断数据类型
	str := "this is string"
	i := 10
	flag := true
	// 调用
	judgeType(str)
	judgeType(i)
	judgeType(flag)
}

// 或者我们可以写一个方法来专门对数据类型进行判断

func judgeType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("当前为 string 类型，value:", v)
	case int:
		fmt.Println("当前为 int 类型，value:", v)
	case bool:
		fmt.Println("当前为 bool 类型，value:", v)
	default:
		fmt.Println("无判断类型，value:", v)
	}

}
