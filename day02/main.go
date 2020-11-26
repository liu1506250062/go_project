package main

import "fmt"

//常量的定义
const pi = 3.1415
const e1 = 2.7182

const (
	statusOk    = 202
	statusError = 404
)

//批量声明常量时候 如果后面的没有赋值 就等于上面的值
const (
	n1 = 100
	n2
	n3
)

//iota 是go语言的常量计数器，只能在常量的表达式中使用。

//PS
//iota在const关键字出现时将被重置为0。
//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用

const (
	a1 = iota //0
	a2        //1
	a3        //2
)


const (
	b1 = iota //0
	b2        //1
	_		  //2
	b3        //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const n5 = iota //0

//多个iota定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

/*定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，
也就是由10变成了1000，也就是十进制的8。）*/
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main() {
		a := 10
        b := &a
        fmt.Printf("a:%d prt:%d \n", a, &a)
        fmt.Printf("b:%p type:%T \n", b, b)
        fmt.Println(&b)
}