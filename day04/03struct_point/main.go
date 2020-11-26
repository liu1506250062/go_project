package main

import "fmt"

//结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {

	var p = new(person)
	fmt.Printf("%T", p)

	//(*p).name = "小王子"
	//(*p).city = "北京"
	//(*p).age = 18

	p.name = "小王子"
	p.city = "北京"
	p.age = 18
	fmt.Printf("%#v\n", p)

	//取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)

	p3.name = "娜扎"
	p3.city = "beijing"
	p3.age = 18
	fmt.Printf("%#v\n",p3)
}
