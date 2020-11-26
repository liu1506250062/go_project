package main

import "fmt"

//为任意类型添加方法

//基于内置的基本类型  造一个我们自己的类型
type MyInt int

func (m MyInt) sayHi()  {
	fmt.Println("hi")
}




func main() {
	var m1 MyInt
	m1 = 100
	m1.sayHi()
}
