package main

import "fmt"

// 循环仍然是以 for 作为关键字
func main() {

	// 案例：计算1～10总和
	// 定义变量
	sum := 0
	i := 10
	// 可以将条件直接跟在for关键字后面
	for i > 0 {
		sum += i
		// 保证循环条件能够结束，不然会进入死循环
		i--
	}
	// 输出结果
	fmt.Println("输出最后的结果为：", sum)
}
