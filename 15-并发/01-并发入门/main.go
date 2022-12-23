package main

import (
	"fmt"
	"time"
)

// 在 go 语言中使用并发的话需要用到 goroutine，使用方式比较简答为 `go 函数名(参数列表)` 即可
func main() {

	// 如果需要调用线程则遵循 goroutine 的调用语法即可
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	// 如果只写到这里，run 当前程序，可能会看见 hello 函数并没有调用十次就结束了，并且数字的顺序不一
	// 原因是当 main 进程结束的时候就会携带着所有在 main 函数中启动的 goroutine 程序回收
	fmt.Println("这里是 main 函数")
	// 如果想要所有的 goroutine 都执行完, 可以添加以下内容让 main 线程在输出完内容后沉睡一会, 给其他线程以执行的时间
	time.Sleep(time.Second)
}

// 当使用 go 关键字调用 hello() 函数的时候，会发现其实执行出来的顺序并不是按照数字顺序的因为线程的调度并不是顺序的而是随机的
func hello(index int) {
	fmt.Println("这里是hello,", index)
}
