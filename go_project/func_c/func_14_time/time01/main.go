package main

import (
	"fmt"
	"time"
)

func dataTime() {
	nows := time.Now()
	fmt.Println(nows.Format("2006-01-02 PM"))
}

func timeParse() {
	// nowtime := time.Now()
	// nowtimes := string(nowtime)
	nowtimes := "2020-04-22"
	timeObj, err := time.Parse("2006-01-02", nowtimes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

func timeF2() {
	now := time.Now() //获取本地时间
	fmt.Println(now)
	timeDate, _ := time.Parse("2006-01-02 15:04:05", "2019-08-04 15:20:20")
	fmt.Println(timeDate)
	//根据字符串加载时区
	localtime, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(localtime)
	// 按照指定的时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-04 15:20:20", localtime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)
}

func main() {
	// now := time.Now()
	// fmt.Println(now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Weekday())
	// fmt.Println(now.Date())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Add(24 * time.Hour))
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }
	// dataTime()
	// timeParse()
	timeF2()
}
