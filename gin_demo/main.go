package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {

	//2 解析模板
	t,err :=  template.ParseFiles("/hello.tmpl")
	if err != nil{
		fmt.Println("1  %v",err)
		return
	}

	// 3 渲染模板
	err = t.Execute(w,"小王子")
	if err != nil{
		fmt.Println("2  %v",err)
		return
	}
}

func main()  {

	http.HandleFunc("/",sayHello)

	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("HTTP server start failed , err : %v", err)
	}
}