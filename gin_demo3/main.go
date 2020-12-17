package main

import (
	"fmt"
	"net/http"
)


//05  4minutes 50 5s

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("qwe")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed ,err :%")
		return
	}
}
