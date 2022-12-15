package main

import (
	"fmt"
	"sync"
)

// 普通的map如果在并发的场景下，会出现读写失效的问题
// 在这种情况下需要用到线程安全的map，线程安全的map不支持普通map的读写方式
// 它有一套自己特有的创建、读写方式
func main() {
	// 创建一个线程安全的map，语法固定不需要使用make
	var mapData sync.Map
	// 添加数据 Store(key, value) 函数
	mapData.Store("k1", "v1")
	mapData.Store("k2", "v2")
	mapData.Store("k3", "v3")
	// 读取数据 Load(key) 函数
	fmt.Println(mapData.Load("k1"))
	// 删除元素
	mapData.Delete("k1")
	// 输出结果
	//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	mapData.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
