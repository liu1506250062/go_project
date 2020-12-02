package main

import (
	"fmt"
	"time"
)


//时间间隔定义

//const (
//	Nanosecond  Duration = 1
//	Microsecond          = 1000 * Nanosecond
//	Millisecond          = 1000 * Microsecond
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)


func main() {

	//Add
	//求两个个小时之后的时间
	now := time.Now()
	later := now.Add(time.Hour * 2)
	fmt.Println(later)

	//Sub
	//求两个时间之间的差值：
	later1 := now.Sub(later)
	fmt.Println(later1)


}
