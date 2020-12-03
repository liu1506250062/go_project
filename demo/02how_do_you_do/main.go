package main

import (
	"fmt"
	"strings"
)

// 统计单词出现次数
func main() {

	s := "how do you do"
	s1 := strings.Split(s, " ")

	m1 := make(map[string]int, 10)
	for _, w := range s1 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
