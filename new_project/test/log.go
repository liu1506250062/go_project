package main

import (
	"fmt"
)

func main() {
	//log.Println("这是一条日之后信息")
	//v := "普普通通"
	//log.Printf("这是一条%s日志\n", v)

	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Println("这是一条日志")

	nums := [5]int{3, 2, 4}

	var numbers = make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == 6 {
				numbers = append(numbers, i)
				numbers = append(numbers, j)
				fmt.Printf("{%d,%d}\n", i, j)
			}
		}
	}
}
