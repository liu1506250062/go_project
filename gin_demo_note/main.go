package main

import (
	"example.com/m/v2/dao"
	"example.com/m/v2/routers"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接

	//绑定模型
	dao.InitMySQL()
	//路由
	r := routers.SetupRpoute()

	//启动项目
	r.Run(":9090")
}
