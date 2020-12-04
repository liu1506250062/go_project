package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//基本请求 Get

func main() {

	url := "https://www.liwenzhou.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	//程序在使用完response后必须关闭回复的主体。
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	fmt.Print(string(body))

}

//要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：

func eg() {
	//client := &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
	//}
	//resp, err := client.Get("http://example.com")
	//// ...
	//req, err := http.NewRequest("GET", "http://example.com", nil)
	//// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//resp, err := client.Do(req)
}
