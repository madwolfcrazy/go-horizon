package helper

import "time"

//Find5MinTimePoint 生成当前时间5分钟 点 的unix 时间戳
func Find5MinTimePoint() int64 {
	t1 := time.Now().Unix()
	return t1 - (t1 % 300)
}

// TodayStartTime 今天开始时间 00:00:00
func TodayStartTime() time.Time {
	tNow := time.Now()
	thisDay := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, time.Local)
	return thisDay
}

// ThisMonday0Time 本周开始时间 周一00:00:00
func ThisMonday0Time() time.Time {
	tNow := time.Now()
	dayOffset := int(time.Monday - tNow.Weekday())
	thisMonday := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, dayOffset)
	return thisMonday
}

//LastWeekMonday0Time 上周周一00:00:00
func LastWeekMonday0Time() time.Time {
	thisMondayTime := ThisMonday0Time()
	return thisMondayTime.AddDate(0, 0, -7)
}

//ThisMonth0Time 本月开始时间
func ThisMonth0Time() time.Time {
	tNow := time.Now()
	thisMonth := time.Date(tNow.Year(), tNow.Month(), 1, 0, 0, 0, 0, time.Local)
	return thisMonth
}

//LastMonth0Time 上月开始时间
func LastMonth0Time() time.Time {
	tNow := time.Now()
	lastMonth := tNow.Month() - 1
	lastYear := tNow.Year()
	if lastMonth < 0 {
		lastMonth = time.October
		lastYear -= 1
	}
	tt := time.Date(lastYear, lastMonth, 1, 0, 0, 0, 0, time.Local)
	return tt
}
