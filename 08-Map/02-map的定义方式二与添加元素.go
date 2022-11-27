package main

import "fmt"

// 上一节我们学会了如何定义map
// 其实map还有另一种常用的定义方式，make关键字(与slice 切片相似)
func main() {
	// 用make定义一个map集合
	// map每次的增长扩容会增加 1，所以出于性能考虑如果已经明确当前的map会增长，那么可以在创建map的时候提前定义好容量空间
	mapList := make(map[string]string)
	// 往map集合中添加元素
	mapList["k1"] = "v1"
	mapList["k2"] = "v2"
	// 输出内容
	fmt.Println("得到的map是：", mapList)

	// 如果需要一个key对应多个value，那么我们可以把value的类型定义成一个切片
	mapValues := make(map[string][]int)
	// 创建一个int切片
	nums := []int{1, 2, 3, 4}
	// 设置map的元素
	mapValues["k1"] = nums
	// 输出结果
	fmt.Println(mapValues)

}
