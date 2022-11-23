package main

import "fmt"

// 语法为 type (别名) 基本类型 或 type (别名) = 基本类型。不能自己想类型

// NewInt 将NewInt定义为int类型
type NewInt int

// Alias 将int取一个别名叫IntAlias
type Alias = int

func main() {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 从输出结果来看, a的类型显示为main包下的NewInt类型

	// 将a2声明为IntAlias类型
	var a2 Alias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
	// 从输出结果来看, a2的类型依然是 int 而不是 Alias. 编译完成时, 不会有 Alias 类型.

}
