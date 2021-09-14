package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse Fail：%v", err)
		return
	}

	//渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Parse Fail：%v", err)
		return
	}
}


//主函数
func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed err:%v", err)
		return
	}
}
