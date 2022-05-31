package main

import (
	"fmt"
	"time"
)

// time 工具类

func main() {

	now := time.Now()
	fmt.Println(now) // 2022-05-20 18:03:30.064221 +0800 CST m=+0.000115292

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("year:%v,month:%v,day:%v,hour:%v,minute:%v,second:%v\n", year, month, day, hour, minute, second)

	// 时间戳 从1970.01.01到现在的一个秒数
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp1) // 毫秒: 1653041230
	fmt.Println(timeStamp2) // 纳秒: 1653041230283016000

	// 将时间戳转换为具体的时间格式
	unix := time.Unix(now.Unix(), 0)
	fmt.Println(unix) // 2022-05-20 18:10:39 +0800 CST

	// 时间间隔  单位纳秒
	time.Sleep(5)               // 默认单位纳秒
	time.Sleep(5 * time.Second) // 默认单位纳秒

	// Add
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)

	// Sub
	sub := now.Sub(later) // 计算两个时间的差距
	fmt.Println(sub)

	// Equal 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
	// Before 如果t代表的时间点在u之前，返回真；否则返回假。
	// After 如果t代表的时间点在u之后，返回真；否则返回假。

	// 定时器  Tick()
	// ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	// for i := range ticker {
	// 	fmt.Println(i) // 每秒都会执行的任务
	// }

	// time.Format函数能够将一个时间对象格式化输出为指定布局的文本表示形式，
	// 需要注意的是 Go 语言中时间格式化的布局不是常见的
	//   Y  - m  - d  H  : M  : S
	// 2006 - 01 - 02 15 : 04 : 05.000
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// parseDemo 解析时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
