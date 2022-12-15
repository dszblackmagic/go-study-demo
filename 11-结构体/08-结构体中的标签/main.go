package main

import (
	"fmt"
	"reflect"
)

// 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记
// 标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它（ORM框架的映射原理与此有关）

type TestTag struct {
	// 标签是跟在字段后面的字符串
	fieldOne   string "The first tag"
	fieldTwo   bool   "The fieldTwo tag"
	fieldThree int    "The fieldThree tag"
}

func main() {

	// 创建结构体实例
	tt := TestTag{"test", true, 10}
	// 通过反射获取内容类型
	of := reflect.TypeOf(tt)
	// 遍历(根据字段数)
	for i := 0; i < of.NumField(); i++ {
		// 输出当前字段
		fmt.Println("当前字段为:", of.Field(i).Name)
		fmt.Println("当前字段类型为:", of.Field(i).Type)
		fmt.Println("当前字段标签为:", of.Field(i).Tag)
		fmt.Println("---------------------")
	}

}
