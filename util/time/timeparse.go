package jtime

/*
时间常量
*/
const (

	//定义每分钟的秒数
	SecondsPerMinute int64 = 60
	//定义每小时的秒数
	SecondsPerHour int64 = SecondsPerMinute * 60
	//定义每天的秒数
	SecondsPerDay int64 = SecondsPerHour * 24
)

/*
时间转换函数
*/
func ResolveTime(seconds int64) (day int, hour int, minute int, second int) {
	day = int(seconds / SecondsPerDay)
	seconds = seconds - int64(day)*SecondsPerDay
	//每小时秒数
	hour = int(seconds / SecondsPerHour)
	seconds = seconds - int64(hour)*SecondsPerHour
	//每分钟秒数
	minute = int(seconds / SecondsPerMinute)
	second = int(seconds - int64(minute)*SecondsPerMinute)
	return
}
