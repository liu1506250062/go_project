package main

import (
	"example.com/m/v2/dao"
	"example.com/m/v2/models"
	"example.com/m/v2/routers"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接 defer 关键字 表示程序运行结束之后调用的方法

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//路由
	r := routers.SetupRpoute()

	//启动项目
	r.Run(":9090")
}
