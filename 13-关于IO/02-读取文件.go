package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 写一个读取文件的案例
func main() {

	// 打开一个本地文件 os.Open 方法来打开文件
	file, err := os.Open("13-关于IO/demo-02.txt")
	// 判断读取是否异常
	if err != nil {
		// 打印错误日志
		log.Println("打开文件异常:", err)
		return
	}
	// 处理文件
	// 处理完文件后要记得关闭资源 要用到 defer
	defer file.Close()
	// 读取文件内容
	// 准备缓冲空间与完整的字符容器切片
	var buff [8]byte
	var container []byte

	for {
		// 调用读取方法，存入 buff, 返回值 n 为读取的字节数
		n, err := file.Read(buff[:])
		// 判断是否读取完
		if err == io.EOF {
			// 读取结束
			log.Println("文件内容读取完毕！")
			break
		}
		// 判断是否读取异常
		if err != nil {
			log.Println("读取异常:", err)
			return
		}
		// 存储数据
		container = append(container, buff[:n]...)
	}
	// 将 byte 数组转为 string 进行输出
	fmt.Println("得到文件中的结果为:", string(container))

}
