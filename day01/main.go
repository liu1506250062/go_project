package main

//导入语句
import "fmt"

//入口函数
//1 基本结构
//2 变量和常量

// 声明变量
// var name string
// var age int
// var isOk bool

//批量声明 (推荐使用驼峰命名)
var (
	name string
	age  int
	isOk bool
	isO  string
)

func main() {
	name = "tom"
	age = 16
	isOk = true

	//在go语言中 非全局变量 定义的变量必须使用 不使用就编译不过去
	fmt.Print("name:", name)   //在终端中输出要打印的内容
	fmt.Printf("age: %d", age) // %d占位符  使用变量的值 来替换占位符
	fmt.Println("isOk", isOk)  //打印完指定的内容之后 会在后面加一个换行符

	//声明变量同时赋值
	var s1 string = "test"
	fmt.Println(s1)

	//短变量声明 只能在函数内部使用
	s2 := "s2"
	fmt.Println(s2)

	//匿名变量

	//注意事项
	/*
		1 函数外的每个语句都必须以关键字开始（var、const、func等）
		2 :=不能使用在函数外。
		3 _多用于占位，表示忽略值。
	*/

}
