package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Gendert string
	Age     int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板
	t, err := template.ParseFiles("hello.tmpl")

	if err != nil {
		fmt.Printf("err : %v\n", err)
		return
	}

	//渲染模板
	u1 := User{
		Name:    "小王子",
		Gendert: "男",
		Age:     23,
	}

	m1 := map[string]interface{}{
		"Name":    "小王子",
		"gendert": "男",
		"Age":     23,
	}

	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{}{
		"m1":        m1,
		"u1":        u1,
		"hobbyList": hobbyList,
	})
}

// 5
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed ,err :%")
		return
	}
}
