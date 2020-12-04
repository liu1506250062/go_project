package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//带参数的get 请求

func main() {

	apiUrl := "http://date.miuu.cn/app/ajax/my_invite.php"
	// URL param
	data := url.Values{}
	data.Set("op", "my_topic")

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
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
