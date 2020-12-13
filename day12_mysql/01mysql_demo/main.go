package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Go连接数据库

func main() {

	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/kwe"

	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                   //dns 格式不正确时候会报错
		fmt.Printf("open %s invalid, err :%v", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err :%v", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
}