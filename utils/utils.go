package utils

import (
	"time"
)

func Time2Milli(time time.Time) int64 {
	return time.UnixMilli()
}

func Time2String(time time.Time) string {
	return time.Format("01/02") + " 星期" + convertWeekday2Num(time.Weekday().String()) + " 14:00-17:00"
}

func convertWeekday2Num(day string) string {
	switch day {
	case "Sunday":
		return "日"
	case "Monday":
		return "一"
	case "Tuesday":
		return "二"
	case "Wednesday":
		return "三"
	case "Thursday":
		return "四"
	case "Friday":
		return "五"
	case "Saturday":
		return "六"
	}
	return ""
}
