package main

import "fmt"

// switch 的条件控制主要是自上而下的，每个 case 的分支都是唯一的
// switch 分支的表达式可以是任意类型，不限于常量
func main() {

	// 定义变量
	score := 60
	grade := "A"
	// 简单拿数值进行匹配举例
	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		// 可以用逗号分开，同时匹配多个条件
		grade = "C"
	default:
		grade = "D"
	}
	// 同时可以采用变量判断来对 case 条件进行控制
	switch {
	case grade == "A":
		fmt.Println("最后的评级是优秀！")
	case grade == "B":
		fmt.Println("最后的评级是良好！")
	case grade == "C":
		fmt.Println("最后的评级是及格，还需要努力哦！")
		// fallthrough 关键字标记后会执行跟在后面的 case 内容，无论条件是否满足
		fallthrough
	case grade == "other":
		fmt.Println("无论怎么样，都需要持续学习！")
	}
}
