package main

import "fmt"

//关闭channel
func main() {

	c := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			c <- i
		}

		//close 关闭一个进程
		close(c)   //如果不关闭，会导致死锁

	}()

	for {

		if data,ok := <-c; ok {
			fmt.Println(data)
		}else {
			break
		}
	}

	fmt.Println("main over")
}


//注意 也不可以向一个已经关闭的channel中发数据 也会导致程序报错 panic