package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ",cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在执行，发送元素：", i, "len(c) = ", len(c), ",cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num:", num)
	}

	fmt.Println("main over")


	//特点
	//当channel已经满，再向李曼写数据，就会阻塞
	//当channel为空，从里面取数据也会阻塞
}
