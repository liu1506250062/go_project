package main

import "fmt"

//类型定义
type MyInt int

//类型别名
type NewInt = int

func main() {

	var a MyInt
	fmt.Printf("type: %T value : %v\n", a, a) //type: main.MyInt   value:0

	var x NewInt
	fmt.Printf("type: %T value:%v\n", x, x) //type:int value:0

}
