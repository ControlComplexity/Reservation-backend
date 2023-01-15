package utils

import "time"

func Time2Milli(time time.Time) int64 {
	return time.UnixMilli()
}
