package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //获取当前时间  2020-12-02 11:19:40.8659622 +0800 CST m=+0.009976201
	fmt.Println(now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//时间戳
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)

	//时间戳转时间
	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)
	year1 := timeObj.Year()     //年
	month1 := timeObj.Month()   //月
	day1 := timeObj.Day()       //日
	hour1 := timeObj.Hour()     //小时
	minute1 := timeObj.Minute() //分钟
	second1 := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year1, month1, day1, hour1, minute1, second1)
}
