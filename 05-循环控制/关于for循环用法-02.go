package main

import "fmt"

// 循环的关键字仍然是 for 不变，但是可以有另外的一种用法如下
func main() {

	// 我们的案例还是取一个比较简单的将数字从 1 加到 10 计算他们的总和
	sum := 0
	// 我们手动创建一个索引来方便退出循环
	i := 0
	// 循环
	for {
		i++
		sum += i
		// 注意在 这里需要一个能够让 for 循环退出的条件，不然的话这种写法如果没有退出的出口就会死循环
		if i > 10 {
			// break 关键字代表的意思就是可以退出循环，结束当前循环的意思
			break
		}
	}
	fmt.Println("得到的计算总和为：", sum)
}
