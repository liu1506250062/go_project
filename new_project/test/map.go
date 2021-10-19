package main

import "fmt"

func main() {
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["李四"] = 100
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["张三"])
	//fmt.Println(scoreMap["小明"])
	//fmt.Printf("Type of a:%T\n", scoreMap)


	//userInfo  := map[string]string{
	//	"username":"沙河",
	//	"password":"123",
	//}
	//fmt.Println(userInfo)


	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 12

	v, ok := scoreMap["ss"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("查无此人")
	}

}
