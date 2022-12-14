package timex

import "time"

func GetDateForAMonthToUnixMilliSecond(y, m int) (startDay, endDay int64) {
	startDayOfMonth, endDayOfMonth := GetDateForAMonth(y, m)

	startDay = startDayOfMonth.UnixMilli()
	endDay = endDayOfMonth.UnixMilli()

	return
}

func GetDateForAMonth(y, m int) (startDay, endDay time.Time) {
	startDay = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
	endDay = time.Date(y, time.Month(m+1), 1, 0, 0, 0, -1, time.UTC)
	return
}

func GetDateForADay(y, m, d int) (startOfDay, endOfDay time.Time) {
	startOfDay = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	endOfDay = time.Date(y, time.Month(m), d+1, 0, 0, 0, -1, time.UTC)
	return
}

func GetDateForADayUnixMillisecond(milli int64) (startDay, endDay int64) {
	y, m, d := time.UnixMilli(milli).Date()
	start, end := GetDateForADay(y, int(m), d)

	startDay = start.UnixMilli()
	endDay = end.UnixMilli()
	return
}
