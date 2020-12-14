package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	id   int
	age  int
	name string
}

//连接数据库
func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/java"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err :%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return

}

//单行数据查询
func queryRowDemo() {
	sqlStr := `select * from users where id = ?`
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("queryRowDemo failed,err :%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err %v\n", err)
	}
	queryRowDemo()

}
