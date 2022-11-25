package main

import "fmt"

// 在Go中有一个关键字 goto 也可以配合来退出循环
// 需要注意的是，goto 后面需要跟着标签来决定它去哪儿
func main() {

	// 一个循环，当i等于6的时候，跳到指定的标签位置
	for i := 0; i < 10; i++ {
		// 进行条件判断
		if i == 6 {
			// 标签在goto后面是必填的，但是它也可以用在break和continue的后面 效果相同，只不过非必填。
			goto End
		}
		fmt.Println("输出当前的内容为:", i)
	}
	// 定义一个标签并且给予命名
End:
	fmt.Println("这里是标签，说明你i的值当前为6")
}
