package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// 同时我们也可以利用 go 语言中存在的 zip 包来处理和压缩/解压相关的内容
func main() {

	// 创建和写入一个压缩文件
	// 创建内容缓冲区
	buf := new(bytes.Buffer)
	// 创建压缩文档
	zp := zip.NewWriter(buf)
	// 在压缩文档内创建文件
	file, err := zp.Create("凳子.txt")
	if err != nil {
		log.Println("创建压缩文件内容失败：", err)
	}
	// 往文件中写数据
	_, err = file.Write([]byte("这是压缩文件中写入的内容.."))
	if err != nil {
		log.Println("写入文件内容失败:", err)
	}
	// 关闭压缩文档
	err = zp.Close()
	if err != nil {
		log.Println("关闭文档失败:", err)
	}
	// 将压缩文档的内容写入成文件
	openFile, err := os.OpenFile("13-关于IO/demo.zip", os.O_WRONLY|os.O_CREATE, 0666)
	// 关闭资源
	defer openFile.Close()
	if err != nil {
		log.Println("创建压缩文件失败:", err)
	}
	// 将内容写入
	buf.WriteTo(openFile)
	// 读取压缩文件中的内容
	readZip()
}

// 写一个函数用来读取压缩文件中的内容
func readZip() {
	// 打开一个 zip 格式的文件
	file, err := zip.OpenReader("13-关于IO/demo.zip")
	if err != nil {
		log.Println("打开文件失败:", err)
	}
	// 记得关闭资源
	defer file.Close()
	// 遍历压缩文件中的内容
	for _, f := range file.File {
		// 查看当前文件的文件名
		fmt.Println("文件名为:", f.Name)
		// 打开文件
		of, err := f.Open()
		if err != nil {
			log.Println("打开压缩内部文件失败:", err)
		}
		_, err = io.CopyN(os.Stdin, of, int64(f.UncompressedSize64))
		if err != nil {
			log.Println("读取文件内容失败:", err)
		}
		fmt.Println()
		// 关闭当前资源
		of.Close()
	}
	// 循环结束提示
	fmt.Println("读取 zip 压缩内容完毕！")
}
