package utils

import (
	"time"
)

const (
	Normal = "2006-01-02 15:04"
	Short  = "2006-01-02"
	Time   = "15:04"
)

func FormatForLayout(timestamp int64, layout string) string {
	if timestamp > 9999999999 {
		timestamp = timestamp / 1000
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(layout)
}

// IsSameDay 判断两个日期是否同一天
func IsSameDay(timestamp1 int64, timestamp2 int64) bool {
	return FormatForLayout(timestamp1, Short) == FormatForLayout(timestamp2, Short)
}

// CalDateGap 计算日期间隔，返回日期列表
func CalDateGap(start int64, end int64) []string {
	var result []string
	result = append(result, FormatForLayout(start, Short))

	for {
		if start > 9999999999 {
			start = start / 1000
		}

		if IsSameDay(start, end) {
			break
		}
		st := time.Unix(start, 0)
		newDt := st.Add(time.Hour * 24)
		start = newDt.Unix()

		result = append(result, FormatForLayout(start, Short))
	}

	return result
}

// CalTimeDiff 计算两个时间小时差
func CalTimeDiff(timestamp1 int64, timestamp2 int64) int64 {
	if timestamp1 > 9999999999 {
		timestamp1 = timestamp1 / 1000
	}
	tm1 := time.Unix(timestamp1, 0)

	if timestamp2 > 9999999999 {
		timestamp2 = timestamp2 / 1000
	}
	tm2 := time.Unix(timestamp2, 0)

	return int64(tm2.Sub(tm1).Hours())
}
