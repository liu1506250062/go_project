package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`  //tag 让字段对应
	Password string `form:"password" json:"password"`
}

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("index.html")
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}

		var u UserInfo
		err := c.ShouldBind(&u) //反射
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/form", func(c *gin.Context) {

		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})

	r.Run(":9090")

}
