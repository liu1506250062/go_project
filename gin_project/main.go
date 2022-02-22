package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

//数据库连接
var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_api?charset=utf8&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	//定义数据库的连接
	//xmlreader
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	r := gin.Default()

	//写入日志
	r.POST("/logAdd", logAdd)

	r.Run(":9090")
}

//定义一个结构体
type Log struct {
	L_message     string `db:"l_message"`
	l_source      int    `db:"l_source"`
	l_type        string `db:"l_type"`
	l_monitor     int    `db:"l_monitor"`
	l_create_time int    `db:"l_create_time"`
	l_status      int    `db:"l_status"`
}

//写日志
func logAdd(c *gin.Context) {

	message := c.PostForm("message")
	source := c.PostForm("source")
	l_type := c.PostForm("type")
	monitor := c.PostForm("monitor")
	if message == "" {
		c.JSON(http.StatusOK, gin.H{"error": "message is not null"})
		return
	}

	r, err := db.Exec("insert into logs(L_message, l_source, l_type,l_monitor,l_create_time,l_status) values(?, ?, ?,?, ?, ?)", message, source, l_type, monitor, time.Now().Unix(), 1)
	if err != nil {
		c.JSON(0, gin.H{"error": err})
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": id})
		return
	}
	c.JSON(0, gin.H{"error": err})
	return
}
