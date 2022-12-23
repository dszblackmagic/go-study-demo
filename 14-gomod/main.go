package main

import (
	"github.com/labstack/echo"
	"net/http"
)

/**
go mod 是 Go 语言中用来方便管理第三方引入的包的工具
go mod 的三种模式：
GO111MODULE=off 禁用 go module，编译时会从 GOPATH 和 vendor 文件夹中查找包；
GO111MODULE=on 启用 go module，编译时会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖；
GO111MODULE=auto（默认值），当项目在 GOPATH/src 目录之外，并且项目根目录有 go.mod 文件时，开启 go module。
*/

// 1. 先在当前目录下使用 go mod init 指令初始化项目(注意目录下不能有中文)
// 2. 使用 go mod tidy 指令自动下载导入相关的依赖包(如果下载不下来可以去配置 GOPROXY的值为国内代理即可[百度])
// 执行以上操作即可完成,对于当前项目下包的管理，并且能看到 go.mod 和 go.sum 两个文件
func main() {

	// 创建一个服务
	e := echo.New()
	// 说明服务接收 get 请求， 第一个参数 ：表示接收的请求露肩，第二个参数：接收到请求后执行的方法
	e.GET("/", func(c echo.Context) error {
		// 成功接收到请求后返回结果为 "Hello World"
		return c.String(http.StatusOK, "Hello, World!")
	})
	// 服务的启动端口为 6666
	// 点击 main 启动后在浏览器访问 `localhost:8888` 即可
	e.Logger.Fatal(e.Start(":8888"))
}

/*
go mod 的其他常用指令如下：
go mod download	下载依赖包到本地（默认为 GOPATH/pkg/mod 目录）
go mod edit	编辑 go.mod 文件
go mod graph	打印模块依赖图
go mod init	初始化当前文件夹，创建 go.mod 文件
go mod tidy	增加缺少的包，删除无用的包
go mod vendor	将依赖复制到 vendor 目录下
go mod verify	校验依赖
go mod why	解释为什么需要依赖
*/
