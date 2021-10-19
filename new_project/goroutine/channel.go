package main

import (
	"fmt"
)

func main() {

	//定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("gorountine 正在运行...")
		c <- 666
	}()

	num := <-c

	fmt.Println(num)
	fmt.Println("main gorountine 结束")
}
