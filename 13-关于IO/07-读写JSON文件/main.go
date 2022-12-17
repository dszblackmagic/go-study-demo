package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Person 符合 Json 文件结构的结构体
type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

// Go 语言中内置了 Json 的读取解析工具可以方便的读取解析写入 Json 文件内容
func main() {

	// 读取文件,
	// os.O_RDWR：打开既可以读取又可以写入文件  os.O_CREATE：如果文件不存在就创建一个新文件
	wrJson, wrErr := os.OpenFile("13-关于IO/person.json", os.O_RDWR|os.O_CREATE, 0666)
	// 关于读取文件的句柄
	rdJson, err := os.Open("13-关于IO/person.json")
	if wrErr != nil {
		log.Println("创建文件失败:", err)
	}
	// 判断异常
	if err != nil {
		log.Println("文件打开失败:", err)
	}
	// 关闭资源
	defer wrJson.Close()
	defer rdJson.Close()
	// 可以利用 json 解析器，将内容写入到 json 文件中
	// 创建实例
	addData := []Person{
		{"李白", 18, []string{"喝酒", "吟诗"}},
		{"杜甫", 19, []string{"忧国", "吟诗"}}}
	// 创建文件输入对象
	encoder := json.NewEncoder(wrJson)
	err = encoder.Encode(&addData)
	if err != nil {
		fmt.Println("写入 json 数据失败:", err)
	}
	// 操作资源 读取 json 数据
	// 利用 go 语言中自带的 json 解码器进行操作
	var persons []Person
	// 用 json 解码器读取文件
	decoder := json.NewDecoder(rdJson)
	// 将内容读取进 persons 这个对象中(传递的是对象的地址)
	err = decoder.Decode(&persons)
	if err != nil {
		log.Println("读取 json 数据失败:", err)
		return
	}
	fmt.Println("读取到的数据为:", persons)

}
