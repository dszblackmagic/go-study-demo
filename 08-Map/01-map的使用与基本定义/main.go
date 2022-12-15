package main

import "fmt"

// map是一个键值对集合，有key与value两个部分组成，每个key对应一个value，可以通过key来快速搜索获得value
// 注意的是 map 是无序的，我们无法决定它的返回顺序
func main() {

	// 定义一个map 语法： var 变量名 = map[keyType]valueType{key:value,...}
	// 初始化一个map
	var mapDemo = map[string]string{"k1": "v1", "k2": "v2"}
	// 输出map
	fmt.Println(mapDemo)
}
