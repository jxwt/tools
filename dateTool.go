package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetNowDate 获取当日日期
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime 获取当前时间
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeFormatToDate 时间转日期
func TimeFormatToDate(d string) string {
	t, _ := time.Parse("2006-01-02 15:04:05", d)
	return t.Format("2006-01-02")
}

// GetTimestampStr 获取当前时间戳字符串
func GetTimestampStr() string {
	return time.Now().Format("20060102150405")
}

// GetTimeNumberByFormat 根据formatStr赋予时间字符串
func GetTimeNumberByFormat(inputTime time.Time, formatStr string) string {
	dateStr := formatStr
	dateStr = strings.Replace(dateStr, "Y", FormatIntPrefixNum(inputTime.Year(), "0", 4), -1)
	dateStr = strings.Replace(dateStr, "m", FormatIntPrefixNum(int(inputTime.Month()), "0", 2), -1)
	dateStr = strings.Replace(dateStr, "d", FormatIntPrefixNum(inputTime.Day(), "0", 2), -1)
	dateStr = strings.Replace(dateStr, "H", FormatIntPrefixNum(inputTime.Hour(), "0", 2), -1)
	dateStr = strings.Replace(dateStr, "i", FormatIntPrefixNum(inputTime.Minute(), "0", 2), -1)
	dateStr = strings.Replace(dateStr, "s", FormatIntPrefixNum(inputTime.Second(), "0", 2), -1)
	return dateStr
}

// GetSubTime 返回时间差 time.Duration
func GetSubTime(start string, end string, format string) time.Duration {
	st, _ := time.Parse(format, start)
	en, _ := time.Parse(format, end)
	return en.Sub(st)
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return getMaxTime(d)
}

func getMaxTime(d time.Time) time.Time {
	loc, _ := time.LoadLocation("Local")
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, loc)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetTimeString 获得天 小时 分钟 秒
func GetTimeString(duration time.Duration) string {
	spanSeconds := int64(duration.Seconds())
	if spanSeconds/(24*60*60) > 0 {
		return fmt.Sprintf("%d天%d小时%d分%d秒", spanSeconds/(24*60*60), spanSeconds%(24*60*60)/(60*60), spanSeconds%(60*60)/60, spanSeconds%60)
	}
	return fmt.Sprintf("%d小时%d分%d秒", spanSeconds%(24*60*60)/(60*60), spanSeconds%(60*60)/60, spanSeconds%60)
}

// BeforeNowLimitSecond 取当前时间的前second秒钟
func BeforeNowLimitSecond(second int) string {
	now := time.Now()
	t := now.Add(-time.Duration(second) * time.Second)
	return t.Format("2006-01-02 15:04:05")
}

// BeforeNowLimitMinute 取当前时间的前minute分钟
func BeforeNowLimitMinute(minute int) string {
	now := time.Now()
	t := now.Add(-time.Duration(minute) * time.Minute)
	return t.Format("2006-01-02 15:04:05")
}

// BeforeNowLimitDay 取当前时间的前day天
func BeforeNowLimitDay(day int) string {
	now := time.Now()
	t := now.Add(-time.Duration(day) * time.Hour * 24)
	return t.Format("2006-01-02 15:04:05")
}

// GetDateTimeMin 获取一天最小时间
func GetDateTimeMin(date string) time.Time {
	if len(date) > 10 {
		date = date[:10]
	}
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return t
}

// GetDateTimeMax 获取一天最大时间
func GetDateTimeMax(date string) time.Time {
	if len(date) > 10 {
		date = date[:10]
	}
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", date+" 23:59:59", time.Local)
	return t
}

// GetTomorrowTimestamp 获取明日时间戳
func GetTomorrowTimestamp() int64 {
	timeLater := time.Now().Add(time.Hour * 24).Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", timeLater, loc)
	return theTime.Unix()
}

// ParseToDateTime 时间字符串转time
func ParseToDateTime(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", d)
	loc, _ := time.LoadLocation("Local")
	t.In(loc)
	return t, err
}

// ParseToDate 日期字符串转time
func ParseToDate(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", d)
	loc, _ := time.LoadLocation("Local")
	t.In(loc)
	return t, err
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	loc, _ := time.LoadLocation("Local")
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, loc)
}

// TimeSubDays t2和t1相差天数
func TimeSubDays(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()
	if hours < 0 {
		return -1
	} else if hours == 0 {
		return 0
	}
	if hours < 24 {
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		}
		return 1
	}
	if (hours/24)-float64(int(hours/24)) == 0 {
		return int(hours / 24)
	}
	return int(hours/24) + 1
}

// GetTimeTypeAndFormat 获取前端的时间类型并且格式化
// eg:传入 '2019' 返回 year,2019-01-01 2020-01-01
// eg:传入 '2019-02' 返回 month,2019-02-01,2019-03-01
// eg:传入 '2019-02-04' 返回 day,2019-02-04,2019-02-05
func GetTimeTypeAndFormat(date string) (dateType, startDate, endDate string) {
	// 筛选类型判断
	if len(date) == 4 {
		dateType = "year"
		date += "-01-01"
	} else if len(date) == 7 {
		dateType = "month"
		date += "-01"
	} else if len(date) == 10 {
		dateType = "day"
	}
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02", date, loc)
	switch dateType {
	case "year":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(1, 0, 0).Format("2006-01-02")
	case "month":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(0, 1, 0).Format("2006-01-02")
	case "day":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(0, 0, 1).Format("2006-01-02")
	}
	return
}

// GetZeroTimeString 获取某一天的0点时间
func GetZeroTimeString(d string) string {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", d, loc)
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, loc)
	return tm.Format("2006-01-02 15:04:05")
}

// GetTimeYearMonthDay 获取传入时间的年月日
func GetTimeYearMonthDay(date string) (year, month, day string) {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	return strconv.Itoa(tm.Year()), strconv.Itoa(int(tm.Month())), strconv.Itoa(tm.Day())
}

// GetHourDiffer 获取相差时间
func GetHourDiffer(startTime, endTime time.Time) int64 {
	var diff int64
	if startTime.Before(endTime) {
		diff = GetSecondDiffer(startTime, endTime)
	} else {
		diff = GetSecondDiffer(endTime, startTime)
	}
	return diff / 3600
}

// GetSecondDiffer 获取相差时间
func GetSecondDiffer(startTime, endTime time.Time) int64 {
	var hour int64
	if startTime.Before(endTime) {
		hour = endTime.Unix() - startTime.Unix()
	} else {
		hour = startTime.Unix() - endTime.Unix()
	}
	return hour
}

// GetFirstDateOfWeek 获取一周的日期
func GetFirstDateOfWeek(begin, max int) []string {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weeks := make([]string, 0)
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	oneWeek := weekStartDate.Format("2006-01-02")
	TimeMonday, _ := time.Parse("2006-01-02", oneWeek)
	for i := begin; i < max; i++ {
		week := TimeMonday.AddDate(0, 0, i)
		weekMonday := week.Format("2006-01-02")
		weeks = append(weeks, weekMonday)
	}
	return weeks
}

// GetDefaultZeroTime 数据库默认时间
func GetDefaultZeroTime() time.Time {
	loc, _ := time.LoadLocation("Local")
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 15:04:05", loc)
	return location
}

// GetDayAndWeek .
func GetDayAndWeek(begin, max int, all []string) ([]string, []string, []string) {
	rr := make([]string, 0)
	rd := make([]string, 0)
	week := make([]string, 0)
	yesterday := time.Now().AddDate(0, 0, -1)
	//一周日期
	weeks := GetFirstDateOfWeek(begin, max)
	weekMap := make(map[string]string, 0)
	weekMap["一"] = weeks[0]
	weekMap["二"] = weeks[1]
	weekMap["三"] = weeks[2]
	weekMap["四"] = weeks[3]
	weekMap["五"] = weeks[4]
	weekMap["六"] = weeks[5]
	weekMap["日"] = weeks[6]
	for _, v := range all {
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02", weekMap[v], loc)
		if t.Before(yesterday) {
			continue
		}
		rr = append(rr, t.Format("01月02日"))
		rd = append(rd, weekMap[v])
		week = append(week, "周"+v)
	}
	return rr, rd, week
}
