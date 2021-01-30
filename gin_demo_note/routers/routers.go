package routers

import (
	"example.com/m/v2/controller"
	"github.com/gin-gonic/gin"
)

func SetupRpoute() *gin.Engine {
	r := gin.Default()

	//告诉 gin 去哪里找模板文件
	r.Static("/static", "static")

	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{

		//添加
		v1Group.POST("/todo", controller.CreateAtodo)

		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTOdoList)

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteAtodo)
	}
	return r
}
