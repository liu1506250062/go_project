package main

func main() {

	//定义一个切片
	//a := [5]int{1,2,3,4,5}
	//s := a[1:3:5]
	//fmt.Printf("%v,%v,%v", s,len(s),cap(s))

	//citySlice := []string{}
	//citySlice = append(citySlice, "北京")
	//citySlice = append(citySlice,"上海","广州","天津")
	//fmt.Printf("%v", citySlice)

	//var a = make([]string ,1 ,3)
	//for i:= 0;i<10;i++{
	//	a = append(a,fmt.Sprintf("%v",i))
	//}
	//fmt.Println(a)






}


func twoSum(nums []int, target int) []int {

	var numbers = make([]int,0)

	for i := 0;i<len(nums);i++{
		for j:=i;j<len(nums);j++{
			if nums[i]+nums[j] == target {
				numbers = append(numbers, i)
				numbers = append(numbers, j)
				return numbers
			}
		}
	}
	return numbers
}

