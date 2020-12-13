package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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

type Users struct {
	id   int
	name string
	age  int
}

// 单行查询
func queryOne(id int) {
	var u1 Users
	// 1 写查询单条的 SQL 语句
	sqlStr := `select * from users where id = ?;`
	//2 执行
	rowObj := db.QueryRow(sqlStr, id)

	//3 拿到结果
	rowObj.Scan(&u1.id, &u1.name, &u1.age) //Scam 里会释放连接  这里一定要调用

	fmt.Printf("u1,%v\n", u1)
}

//多行查询
func queryNorelco(n int) {
	sqlStr := `select * from users  where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec $s query faild, err:%v\n", sqlStr, err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var u1 Users
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scam faild, err %v\n", err)
		}
		fmt.Printf("u1 : %#v\n", u1)
	}

}

//插入数据
func insert() {
	sqlStr := `insert into users(name ,age) values("test",123)`
	ret, err := db.Exec(sqlStr)

	if err != nil {
		fmt.Printf("scam faild, err %v\n", err)
		return
	}

	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id faild, err %#v\n", err)
		return
	}

	fmt.Println("id :", id)
}

//更新数据
func updateRow(newAge, id int) {
	sqlStr := `update users set age = ? where id = ?;`
	n, err := db.Exec(sqlStr, id, newAge)
	if err != nil {
		fmt.Printf("scam faild, err %v\n", err)
		return
	}
	fmt.Println("n :", n)
}

//删除
func deleteRow(id int) {
	sqlStr := `delete from users where id = ?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("deleteRow faild, err %v\n", err)
		return
	}
	fmt.Println("n :", ret)
}

//预处理插入多条数据
func prepareInsert() {
	sqlStr := `insert into users(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepareInsert faild, err %v\n", err)
	}

	defer stmt.Close()

	//后续只要拿到是说stmt 去执行一些操作
	var m = map[string]int{
		"cecshi": 12,
		"cess":   45,
		"rtete":  45,
		"ewqe":  45,
	}
	for k,v := range m{
		stmt.Exec(k,v)
	}

}

// Go连接数据库
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err %v\n", err)
	}
	fmt.Println("success")

	//queryOne(2)   //查询单行数据
	//queryNorelco(0) //查询多行数据
	//insert()   //添加数据
	//updateRow(2,1231)   //更新数据
	//deleteRow(2) //删除数据
	prepareInsert() //批量处理
}
