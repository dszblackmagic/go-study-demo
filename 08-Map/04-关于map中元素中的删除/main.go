package main

import "fmt"

// 我们讲述了map的创建、元素的添加、以及map的遍历
// 那么这个案例我们来讲述一下如何对map进行删除，需要使用到 delete 这个关键函数
func main() {

	// 定义一个map集合
	mapData := make(map[string]string)
	// 添加元素
	mapData["k1"] = "这是v1"
	mapData["k2"] = "这是v2"
	mapData["k3"] = "这是v3"
	// 先输出一次
	fmt.Println("当前的map数据为：", mapData)
	// 我们知道我们可以根据key 找到对应的 value 所以我们想要删除map中的某个元素，我们就通过key 来进行删除即可
	// 我们要删除 mapData 中key为'k1'所对应的元素
	delete(mapData, "k1")
	// 输出集合
	fmt.Println("删除后的map数据为：", mapData)

}
