package main

import "fmt"

func main() {

	// 除了使用&来指向指针，也可以用另一种方式创建指针
	// 语法采用 new() -> 这时候的变量就等于创建了一个指针
	str := new(string)
	// 指向默认开辟的指针地址
	fmt.Println(str)
	// 赋值后再看看
	*str = "现在就一起let's go"
	fmt.Println(*str)
}
