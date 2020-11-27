package main

import "fmt"

//方法的定义与实例
type Person struct {
	name string
	age  int8
}

//NewPerson 是一个Person的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//定义方法 方法的定义 比函数前多了一个 接受者
func (p Person) Dream() {
	fmt.Printf("%s的dream是XX", p.name)
}

//修改年龄 指针接受者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

//修改年龄 值接受者
func (p Person) setAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("沙河", int8(12))
	//p1.Dream()

	fmt.Println(p1.age) //12
	p1.SetAge(int8(23))
	fmt.Println(p1.age) //23   如果想要改变值 必选传指针
	p1.setAge2(int8(34))
	fmt.Println(p1.age) //23
}

//Ps::
//函数不属于任何类型
//方法属于具体类型

//什么时候应该用指针类型
//1 需要修改接受者中的值
//2 接受者是拷贝代价比较大的对象
//3 保证一致性 如果某个方法使用了指针接受者 那么其他的方法也应该使用指针接受者