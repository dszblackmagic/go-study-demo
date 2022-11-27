package main

import "fmt"

// 同样的对于map这个类型里面的数据我们也有需要遍历的场景
// 那么我们可以通过for range 来遍历当前的map集合
func main() {

	// 定义并且初始化map对象
	mapList := map[string]int{"k1": 10, "k2": 20, "k3": 30}
	// 开始遍历
	for key, value := range mapList {
		// 输出内容
		fmt.Println("遍历的key为：", key, " 当前key所对应的value为：", value)
	}
}
