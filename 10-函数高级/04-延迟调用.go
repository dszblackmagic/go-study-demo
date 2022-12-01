package main

import (
	"fmt"
	"log"
)

// 延迟调用：在 go 语言中延迟调用简单理解来说就是会最后执行。可以用来处理一些资源释放问题（文件资源、数据库连接、锁）
func main() {

	// 如果有学习过Java语言也可以将 defer 理解成 finally 块比较类似
	// 看以下一个案例
	fmt.Println("输出一")
	// 用 defer 输出一个日志
	defer log.Println("打印一个 defer 日志")
	// 虽然下面代码在defer之后，但是从输出结果来看的话，会在defer之前输出，因为defer标记是延迟调用会在最后输出
	fmt.Println("输出二")
}
