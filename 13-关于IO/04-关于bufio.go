package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 在我们写入内容的过程中，我们都是自己去定义缓冲区(buf),然后每次将内容暂存在缓冲区中由我们自己进行获取
// 实际上在 go 语言中则有一个包为 bufio 帮我们处理好了这些
func main() {
	// 调用方法先写入内容、后读取
	wr()
	rd()
}

/**
关于系统 os 处理文件的一些模式：
os.O_WRONLY	只写
os.O_CREATE	创建文件
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加
*/
// 写一个关于写入内容的方法
func wr() {

	// 我们可以打开一个文件往里面追加内容
	// 第一个参数为创建（打开）的文件名， 第二个参数为 处理方式的分位符(如果不存在就创建、赋予写的权限)
	// 第三个参数则是对应 unix中的文件权限 比如0777：-rwxrwxrwx，创建了一个普通文件，所有人拥有所有的读、写、执行权限
	file, err := os.OpenFile("13-关于IO/demo-04.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		// 如果文件打开异常则输出日志
		log.Println("文件创建异常:", err)
	}
	// 同样的在使用的时候要记得关闭资源
	defer file.Close()
	// 利用 bufio 获取文件写入对象更方便我们操作
	writer := bufio.NewWriter(file)
	// 可以利用 bufio 生成的写入对象对文件进行操作
	for i := 0; i < 5; i++ {
		// 可以随意写入自己需要的内容
		writer.WriteString("this is string " + "\n")
	}
	// 刷新缓冲区将内容保存
	writer.Flush()
}

// 同样的我们可以再写一个读取的方法来读取我们之前写入的内容
func rd() {
	// 同样的打开文件
	file, err := os.Open("13-关于IO/demo-04.txt")
	// 判断文件读取情况
	if err != nil {
		log.Println("文件读取异常:", err)
	}
	// 处理文件完需要关闭资源
	defer file.Close()
	// 获取文件的读取类
	reader := bufio.NewReader(file)
	for {
		// 利用 reader 读取数据
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			log.Println("文件已经读取完毕！")
			break
		}
		if err != nil {
			log.Println("文件读取异常:", err)
			return
		}
		// 如果以上校验没有问题的话则输出文件内容
		fmt.Println(string(line))
	}
}
