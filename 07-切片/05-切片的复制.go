package main

import "fmt"

// 切片之间可以用copy() 来进行复制
// 如果两个切片不一样大小，那么复制的时候则会按照小的那个进行内容复制
// 注意：进行切片复制的两个切片必须同类型
func main() {

	// 定义两个切片
	sourceSlice := []int{1, 2, 3}
	targetSlice := []int{4, 5, 6, 7, 8}
	// 进行内容的copy
	// 第一个参数为目标切片， 第二个参数数据源切片
	copy(targetSlice, sourceSlice)
	// 输出复制后的切片
	// 可以看到结果除了前三个内容发生变化外其他原本的内容没有变化，这就和切片复制时候的规则有关（按照小的切片元素来）
	fmt.Println("复制后的切片内容为：", targetSlice)
}
