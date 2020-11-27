package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性 与 JSon序列化

//如果一个Go语言中定义的标识符首字母是大写的，那么就是对外可见
//如果一个结构体中的字段首字母是大写的 那么他就是对外可见的

type student struct {
	ID   int
	Name string
}

type class struct {
	Title     string    `json:"title"`
	Studetent []student `json:"student" db:"student"`
}

// student 的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {

	c1 := class{
		Title:     "测试",
		Studetent: make([]student, 0, 20), //切片使用是需要初始化的
	}

	for i := 0; i < 10; i++ {

		//造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Studetent = append(c1.Studetent, tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	//JSON 序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json error", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//json 反序列化
	//jsonStr := `{"Title":"测试","Studetent":[{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}`
	//var c2 class
	//err = json.Unmarshal([]byte(jsonStr), &c2)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Printf("%#v\n", c2)
}
