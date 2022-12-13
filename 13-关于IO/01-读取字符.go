package main

import (
	"io"
	"log"
	"strings"
)

// 关于IO的应用常见的方式多是为'读取(Reader)'，'写入(Writer)'，'关闭(Closer)'
// 读取的内容可以是任何可以解析的东西，包括：文字、文件等等...

/*
*
案例： 用 Reader 来读取文字
*/
func main() {

	// 创建一个字符串变量
	str := "this is string"
	// 可以创建一个字符串读取器 NewReader() 可以读取字符串(返回值是 *Reader)
	reader := strings.NewReader(str)
	// 每次读取的字节可以将它装起来用 byte 切片
	// 每次读取过程中我们装载 4 个字节
	container := make([]byte, 4)
	// 用 for 循环来进行重复读取
	for {
		// 用 Reader 中的 Read 方法对内容进行读取
		// 得到返回值 n : 表示读取到的字节数， err 则表示异常
		n, err := reader.Read(container)
		// 判断异常是否为空, 如果有异常值那么则需要处理
		if err != nil {
			// io.EOF 表示的是内容已经读取完毕
			if err == io.EOF {
				// 如果内容已经读取完毕(也就是当前读取到的字节为0)，那么我们就需要结束当前的循环
				log.Printf("内容已经读取完毕, %d", n)
				break
			}
			log.Println("其他异常:", err)
		}
		// 如果没有进入异常，则正常输出内容
		// 输出内容的时候 可以用切片只截取当前读到的字节数
		log.Printf("当前读取到字节数为[%d], 内容为: %s", n, string(container[:n]))
	}
}
