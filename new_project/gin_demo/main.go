package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}

func main() {

	//创建一个默认的路由引擎
	r := gin.Default()

	//指定用户使用get方式
	//r.GET("/hello", sayhello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.Run(":9090")
}
