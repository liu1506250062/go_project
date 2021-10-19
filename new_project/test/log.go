package main

import "log"

func main() {
	//log.Println("这是一条日之后信息")
	//v := "普普通通"
	//log.Printf("这是一条%s日志\n", v)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条日志")

}
