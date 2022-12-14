package main

import (
	"fmt"
	"log"
	"os"
)

// 一些在 go 语言中常用的快速文件处理工具方法(在1.16之前是 ioutil, 后来被废弃改到 os、io/fs 包下)
// ReadAll(从某个源读取数据)、ReadFile（读取文件内容）、WriteFile（将数据写入文件）、ReadDir（获取目录） ..
func main() {

	// 快速读取文件(返回字节数组)
	content, err := os.ReadFile("13-关于IO/demo-03.txt")
	if err != nil {
		log.Println("出现异常:", err)
	}
	// 输出读取结果
	fmt.Println(string(content))
	fmt.Println("---------------")
	// 快速写入内容到文件
	WrErr := os.WriteFile("13-关于IO/demo-03.txt", []byte("测试写入内容..."), 0666)
	if WrErr != nil {
		log.Println("写入内容异常:", WrErr)
	}
}
