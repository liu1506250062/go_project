package main

import "fmt"

type Person struct {
	Name   string
	Gender string
	Age    int
	Address
	Email
}

type Address struct {
	Province   string
	City       string
	Updatetime string
}

type Email struct {
	Addr       string
	Updatetime string
}

func main() {

	p := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    23,
		Address: Address{
			Province:   "山东",
			City:       "威海",
			Updatetime: "2020-03-31",
		},
		Email: Email{
			Addr:       "222@22,com",
			Updatetime: "2020-02-12",
		},
	}

	fmt.Printf("%#v\n", p)
	fmt.Println(p.Address.Updatetime)   //包含同名的字段 不能直接访问
}
