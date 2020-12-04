package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	url := "http://date.miuu.cn/app/ajax/my_invite.php"

	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json

	contentType := "application/json"
	data := `{"op":"my_topic","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}

	fmt.Println(string(b))
}
