package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) //这个地方不能用 冒号等于  因为 冒号等于就是在方法内申明   而这个地方 是为了申明一个全局的
	if err != nil {
		return
	}
	return DB.DB().Ping() //返回一个错误
}

func Close() (err error) {
	DB.Close()
	return
}
