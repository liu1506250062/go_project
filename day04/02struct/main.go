package main

import "fmt"

//定义结构体
type Person struct {
	name string
	age  int8
	city string
}

func main() {

	var p Person
	p.name = "沙河"
	p.city = "北京"
	p.age = 18

	fmt.Printf("p= %#v\n", p)
	fmt.Println(p.name)
	fmt.Println(p.city)
	fmt.Println(p.age)

	//匿名结构体 （只用一次的结构体）
	var user struct {
		name    string
		married bool
	}

	user.name = "小王子"
	user.married = false
	fmt.Printf("%#v\n",user)

}
