package main

import (
	"fmt"
	"strconv"
)

//strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。

//Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()

//Format系列函数实现了将给定类型数据格式化为string类型数据的功能。



func main() {

	//Atoi()函数用于将字符串类型的整数转换为int类型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Printf("type : %T, value : %#v\n", i1, i1)
	}


	//Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下
	s2 := 100
	i2 := strconv.Itoa(s2)
	fmt.Printf("type : %T, value : %#v\n", i2, i2)





}
