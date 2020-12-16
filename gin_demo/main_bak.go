package main

import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello golang!")
}

func main() {
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf(" HTTP fail ： %v", err)
	}
}
