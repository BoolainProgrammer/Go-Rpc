package util

import "time"

// 时间处理类

const (
	TIMESTR = "20060102150405"
)

func TimeNow(format string) string {
	return time.Now().Format(format)
}