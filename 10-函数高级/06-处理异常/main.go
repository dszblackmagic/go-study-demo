package main

import (
	"fmt"
	"log"
)

// Go语言中使用 panic 抛出错误，recover 捕获错误。
func main() {
	// 注意一：在 defer 中使用 recover() 获取异常才会捕获到,不然panic会结束执行其之后的普通代码内容
	// 注意二：defer的处理定义必须在 panic 之前, 不然无法生效
	// 定义一个处理异常的匿名函数
	defer func() {
		// 用recover() 捕获异常，并且当异常内容不为空的时候
		if errMsg := recover(); errMsg != nil {
			// 处理异常(打印异常日志信息)
			log.Printf("有异常信息为：{%s} \n", errMsg)
		}
	}()

	// Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
	panic("this is error msg!")
	// 在异常之后的代码会被终止执行
	fmt.Println("这里是异常之后的代码")
}
