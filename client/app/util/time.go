package util

import "time"

// 时间处理类

const (
	TIMESTR = "20060102150405"
)

func TimeNow(format string) string {
	return time.Now().Format(format)
}
//TimeNowStr
//获得年-月-日 时：分：秒时间格式字符串
func TimeNowStr() string {
	t := time.Now()
	timestamp := t.Unix()
	tm := time.Unix(timestamp, 0)
	nowTime := tm.Format("2006-01-02 15:04:05")
	return nowTime
}