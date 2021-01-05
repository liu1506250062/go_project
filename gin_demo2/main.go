package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取参数

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的 query string 参数
		//name := c.Query("query")
		//name, ok := c.GetQuery("query")        // 取不到第二个参数 就返回false

		name := c.DefaultQuery("query", "123") //带默认值的传参
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})

	})

	r.Run(":9090")
}
