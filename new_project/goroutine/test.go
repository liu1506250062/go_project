package main

import (
	"fmt"
	"time"
)

func main() {

	//用go承载一个形参为0
	go func() {

		defer fmt.Println("A.defer ")

		func() {
			defer fmt.Println("B.defer ")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
