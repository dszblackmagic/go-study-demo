package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 利用我们所学习的知识，将 demo-02 中的内容追加到 demo-03.txt 中
func main() {

	// 打开对应的文件
	input, err := os.Open("13-关于IO/demo-02.txt")
	// 设置输入文件内容为 写、追加 以及相关的权限
	output, _ := os.OpenFile("13-关于IO/demo-04.txt", os.O_WRONLY|os.O_APPEND, 0666)
	// 关闭资源
	defer input.Close()
	defer output.Close()
	// 判断打开文件
	if err != nil {
		log.Println("打开文件失败:", err)
	}
	// 读取文件内容
	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	for {
		// 读取文件内容
		line, _, readErr := reader.ReadLine()
		if readErr == io.EOF {
			log.Printf("文件读取完毕！")
			break
		}
		// 开始将内容写入到文件中(可以忽略写入的字节数)
		_, wrErr := writer.WriteString(string(line))
		if wrErr != nil {
			log.Println("写入文件异常:", wrErr)
			break
		}
		writer.Flush()
	}
	// 提示
	fmt.Println("复制追加完成！")

}
