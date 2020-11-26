package main

import "fmt"

//结构体初始化

type person struct {
	name, city string
	age        int8
}

func main() {

	//1 键值对初始化
	p4 := person{
		name: "小王子",
		city: "北京",
		age:  14,
	}

	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}

	fmt.Printf("%#v\n", p5)

	//2 值得列表进行初始化

	p6 := person{
		"小王子",
		"北京",
		28,
	}

	fmt.Printf("%#v\n", p6)

}
