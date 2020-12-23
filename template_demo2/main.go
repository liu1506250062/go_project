package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main() {

	//创建一个理由引擎
	r := gin.Default() //返回一个默认的路由引擎

	//执行用户使用GET 请求的访问/hello 时候 ，执行 sayhello 函数
	r.GET("/hello", sayHello)

	//启动服务
	r.Run(":9090")

}
