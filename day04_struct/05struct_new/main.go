package main

import "fmt"

//结构体的构造函数  构造一个结构体实例的函数

type person struct {
	name, city string
	age        int8
}

//返回数据比较大的时候 一般返回结构体指针
func NewPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p1 := NewPerson("小王子", "北京", 13)
	fmt.Printf("type:%T  value:%#v\n", p1, p1)
}
