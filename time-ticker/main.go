package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func main() {
	t := time.NewTicker(5 * time.Second)
	runNum := 0
	for {
		select {
		case <-t.C:
			runNum++
			fmt.Println("第" + cast.ToString(runNum) + "次开始执行")
			fmt.Println("开始时间:" + GetNowTime())
			time.Sleep(3 * time.Second)
			/*
				第1次开始执行
				开始时间:2022-08-03 15:51:04
				第1次执行完成
				完成时间:2022-08-03 15:51:07
				第2次开始执行
				开始时间:2022-08-03 15:51:09
				第2次执行完成
				完成时间:2022-08-03 15:51:12
				第3次开始执行
				开始时间:2022-08-03 15:51:14
				第3次执行完成
				完成时间:2022-08-03 15:51:17
				第4次开始执行
				开始时间:2022-08-03 15:51:19
				第4次执行完成
				完成时间:2022-08-03 15:51:22

			*/

			//time.Sleep(6 * time.Second)
			/*运行结果
			第1次开始执行
			开始时间:2022-08-03 14:26:13
			第1次执行完成
			完成时间:2022-08-03 14:26:19
			第2次开始执行
			开始时间:2022-08-03 14:26:19
			第2次执行完成
			完成时间:2022-08-03 14:26:25
			第3次开始执行
			开始时间:2022-08-03 14:26:25
			*/

			fmt.Println("第" + cast.ToString(runNum) + "次执行完成")
			fmt.Println("完成时间:" + GetNowTime())
		}
	}

}

// 返回当前时间
func GetNowTime() string {
	cstZone := time.FixedZone("CST", 8*3600) // 东八
	now := time.Now().In(cstZone)
	return TimeFormat("Y-m-d H:i:s", now)
}

// 时间转日期
func TimeFormat(format string, t time.Time) string {
	if format != time.RFC3339 {
		format = strings.Replace(format, "Y", "2006", -1)
		format = strings.Replace(format, "m", "01", -1)
		format = strings.Replace(format, "d", "02", -1)
		format = strings.Replace(format, "H", "15", -1)
		format = strings.Replace(format, "i", "04", -1)
		format = strings.Replace(format, "s", "05", -1)
	}
	if t.IsZero() {
		return ""
	}
	return t.Format(format)
}
