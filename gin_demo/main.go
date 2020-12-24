package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	//r.LoadHTMLFiles("templates/index.tmpl")   //模板解析
	r.LoadHTMLGlob(".templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		//http 请求
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{    //模板渲染
			"title":"liwenzhou.com",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		//http 请求
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{    //模板渲染
			"title":"users/index",
		})
	})

	r.Run(":9090")   //启动server
}
