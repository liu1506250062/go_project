package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

//TODO Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) //这个地方不能用 冒号等于  因为 冒号等于就是在方法内申明   而这个地方 是为了申明一个全局的
	if err != nil {
		return
	}
	return DB.DB().Ping()   //返回一个错误
}

func main() {

	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}


	//绑定模型

	r := gin.Default()

	//告诉 gin 去哪里找模板文件
	r.Static("/static", "static")

	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {})
		{

		}

		//查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})

		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})

		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	r.Run(":9090")

}


// 09  44