package main

import "fmt"

//嵌套结构体

type Person struct {
	Name   string
	Gender string
	Age    int
	//Address Address
	Address //
}

type Address struct {
	Province string
	City     string
}

func main() {

	p := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    34,
		Address: Address{
			Province: "山东",
			City:     "济南",
		},
	}
	fmt.Printf("%#v\n", p)
	fmt.Println(p.Address.Province) // 通过嵌套的匿名结构体字段访问内部字段
	fmt.Println(p.Province)         //直接访问匿名结构体中的字段
}
