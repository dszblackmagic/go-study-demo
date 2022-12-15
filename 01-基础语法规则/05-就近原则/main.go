package main

import "fmt"

var c = "这是全局变量"

func main() {

	c := "这是局部变量"
	fmt.Println("在main中会采用的变量是:", c)
	demo()
}

func demo() {
	fmt.Println("如果没有局部变量影响: ", c)
}
