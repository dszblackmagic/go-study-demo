package main

import (
	"errors"
	"fmt"
	"log"
)

// 除了直接使用panic 异常，我们还可以通过标准库中提供的 errors 来自己创建异常
// errors.New() 可以用来创建一个 error 对象
func main() {
	// 添加错误处理的预防
	defer func() {
		// 直接捕获异常进行输出
		log.Println(recover())
	}()
	// 调用函数
	i, err := division(10, 0)
	if err != nil {
		// 如果err不为nil说明有异常返回
		panic(err)
	}
	// 正常输出结果
	fmt.Println("得到两数计算的结果为:", i)
}

// 假设场景如下, 两数相除，除数不能为0, 如果为 0 则抛出异常（这个异常[error]我们自己定义）
func division(x, y int) (int, error) {
	// 判断除数是否为0
	if y == 0 {
		return 0, errors.New("error:除数不能为零！")
	}
	return x / y, nil
}
