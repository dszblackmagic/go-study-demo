package main

import (
	"log"
	"os"
)

// 前面案例我们了解了如何读取文件，那么后面我们要将内容写入到文件中
func main() {

	// 写入的方法在 io.Writer()
	// 过程与读取内容相似, 需要先从缓冲区读取数据，然后再将读取到的数据写入到目标中
	// 创建一个文件
	file, err := os.Create("13-关于IO/demo-03.txt")
	// 判断是否异常
	if err != nil {
		log.Println("创建文件异常:", err)
	}
	// 最后要关闭资源服务
	defer file.Close()
	// 创建一段内容
	str := "let's go study Writer File"
	// 将内容写入文件，从下面函数名可以看出来，这是在文件中写入 string
	file.WriteString(str)
	// 也可以写入字节数组
	file.Write([]byte(" this is bytes!"))
}
