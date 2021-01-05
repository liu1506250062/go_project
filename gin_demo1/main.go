package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//通过gin框架 返回 json

func main() {
	r := gin.Default()


	//方法1 使用map
	r.GET("/json", func(c *gin.Context) {

		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     14,
		//}

		data := gin.H{
			"name": "小王子", "message": "hello world", "age": 14,
		}

		c.JSON(http.StatusOK, data)
	})


	//方法2 结构提 灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name string `json:"name"`
		Message string
		Age int
	}

	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name: "小王子", 
			Message: "Hello GoLang",
			Age: 12,
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
