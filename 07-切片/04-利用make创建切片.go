package main

import "fmt"

// 上述过程中我们了解来切片的扩容方式以及扩容的过程
// 实际上我们在创建切片的时候就可以来决定切片的大小和预分配空间的大小（也就是容量）
// 我们通过预分配的方式降低了切片进行自动扩容，更好的提升了程序的效率
func main() {

	// 用make定义一个切片
	// 用make创建的这个切片的解析为：string类型的切片, 默认元素2个, 预分配空间为10
	sliceDemo := make([]string, 2, 10)
	// 输出当前元素大小与空间进行验证
	fmt.Println("当前元素大小为：", len(sliceDemo), "当前元素空间容量为：", cap(sliceDemo))
}
