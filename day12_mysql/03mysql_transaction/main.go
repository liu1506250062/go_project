package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	id   int
	name string
	age  int
}

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {

	//数据库信息
	dsn := "root:root@tcp(localhost:3306)/java?charset=utf8mb4&parseTime=True"

	// 连接数据库
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  //dns 格式不正确时候会报错
		fmt.Printf("open %s invalid, err :%v", dsn, err)
		return err
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err :%v", dsn, err)
		return err
	}

	//设置数据库最大连接数
	//db.SetMaxOpenConns(10)

	//设置最大空闲连接数
	//db.SetConnMaxIdleTime(5)

	return nil

}

func transactionDemo() {

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("transactionDemo Begin, err :%v", err)
		return
	}

	sqlStr1 := `update users set age = age-1 where id = 1`
	sqlStr2 := `update users set age = age-1 where id = 1`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()   //事务回滚
		fmt.Printf("sqlStr1 , err :%v", err)
		return
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("sqlStr2 , err :%v", err)
		return
	}

	err = tx.Commit()   //提交事务
	if err != nil {
		fmt.Printf("Commit , err :%v", err)
	}

	fmt.Println("111 ok")

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err %v\n", err)
	}
	fmt.Println("success")

	transactionDemo()
}
