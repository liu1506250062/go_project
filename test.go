package main

import "fmt"

func main() {

}

func f1() {
	fmt.Println("hello world")
}

func f2(name string) {
	fmt.Println("hello %s\n", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}
//test