package main

import (
	"fmt"
	"time"
)


//从gogrountine
func  newTask()  {
	i:= 0
	for {
		i++
		fmt.Printf("new Gorountine.i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}

//主Gorountine
func main() {
	//创建一个go程 执行newTask函数
	go newTask()

	i := 0
	for  {
		i++
		fmt.Printf("Main Gorountine.i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}
