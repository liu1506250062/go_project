package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayhello(w http.ResponseWriter, r *http.Request) {

	//定义模板

	//打开模板
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("Template Failed %v", err)
		return
	}

	u1 := User{
		Name:   "小王子1",
		Gender: "男",
		Age:    12,
	}

	m1 := map[string]interface{}{
		"Name":   "小王子2",
		"Gender": "男",
		"Age":    "12",
	}

	hobbyList := []string{
		"足球",
		"篮球",
		"双色球",
	}

	//渲染模板
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http Failed %v", err)
		return
	}
}
