package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {

	//定义一个函数 kua
	//要么只有一个返回值  要么有两个返回值 第二个返回值 必须是 error 类型
	k := func(name string) (string, error) {
		return name + "123", nil
	}

	//定义模板
	t := template.New("f.tmpl")

	//告诉模板 有一个自定义函数
	t.Funcs(template.FuncMap{
		"kua": k,
	})

	//解析模板
	_, err := t.ParseFiles("f.tmpl") // 创建一个名字是f的模板对象  名字一定要和模板名对应
	if err != nil {
		fmt.Printf("parse template is faild, err : %v", err)
		return
	}
	name := "小王子"

	//渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {

	//定义模板

	//解析模板
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl") //同时解析多个文件 要把被包含的写在后边

	if err != nil {
		fmt.Printf("faild: %v", err)
		return
	}
	name := "小王子"

	//渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		fmt.Printf("Http Server not start  %v", err)
		return
	}
}

