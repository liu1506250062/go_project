package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {

	//连接数据库
	r := gin.Default()

	//告诉 gin 去哪里找模板文件
	r.Static("/static", "static")

	//告诉gin框架去哪里找模板文件
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":9090")

	//v1




}
