package main

import "fmt"

func main() {

	// 设置元素数量为1000
	const elementCount = 1000

	// 预分配足够多的元素切片
	sourceData := make([]int, elementCount)

	// 以遍历的方式将切片中的内容进行复制
	for i := 0; i < elementCount; i++ {
		sourceData[i] = i
	}

	// 引用切片数据
	// 引用并不是复制，只是引用地址
	refData := sourceData

	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, sourceData)

	// 修改原始数据的第一个元素
	sourceData[0] = 999

	// 打印引用切片的第一个元素
	// 因为是引用所以在源数据发生改变的情况下，引用数据也发生改变
	fmt.Println("引用切片的第一元素为：", refData[0])

	// 打印复制切片的第一个和最后一个元素
	// 复制的切片本身具备数据，所以源数据发生改变但是复制切片仍然不变
	fmt.Println(copyData[0], copyData[elementCount-1])

	// 输出原本copyData的前五位数
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	fmt.Println("\n--------分割线------------")
	// 复制原始数据从4到6(不包含)
	// 实际上就是将原数据中的[4,5]两个元素复制到copyData中,因为只有两个元素，所以只改变了copyData的第一第二元素，其他元素仍然不变
	copy(copyData, sourceData[4:6])

	// 输出结果得以验证
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
