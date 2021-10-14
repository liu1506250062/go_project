package main

import "fmt"

func main() {

	//数组定义
	//var testArray [3]int
	//var numArray = [3]int{1, 2}
	//var cityArray = [3]string{"北京", "上海", "深圳"}

	//fmt.Println(testArray)
	//fmt.Println(numArray)
	//fmt.Println(cityArray)

	//遍历数组
	//for i := 0; i < len(numArray); i++ {
	//	fmt.Printf("val = %d\n", numArray[i])
	//}

	//for k, v := range numArray {
	//	fmt.Printf("a := %d=>%d\n", k,v)
	//}

	//二维数组
	//a := [3][2]string{
	//	{"北京", "上海"},
	//	{"北京1", "上海1"},
	//	{"北京2", "上海2"},
	//}
	//
	////遍历二维数组
	//for _, v1 := range a {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%s", v2)
	//	}
	//	fmt.Println()
	//}

	//a := [5]int{1, 3, 4, 6, 7}
	//
	//count := 0
	//for _, v := range a {
	//	count += v
	//}
	//fmt.Printf("count = %d", count)

	//获取数组中 两个元素之和等于8的下标
	array1 := [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(array1); i++ {
		for j := i; j < len(array1); j++ {
			if array1[i]+array1[j] == 8 {
				fmt.Printf("{%d,%d}\n", i, j)
			}
		}
	}








}
