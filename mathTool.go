package tools

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// IntToStringLength .
func IntToStringLength(num int, length int) string {
	s := strconv.Itoa(num)
	for i := length; i < len(s); i++ {
		s = "0" + s
	}
	return s
}

// RandRangeIntNum 产生min-max的随机数
func RandRangeIntNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}

// IntToHex int转十六进制
func IntToHex(num int) string {
	//base,_ := strconv.Atoi(strconv.FormatInt(num,10))
	//return strconv.FormatInt(base, 16)
	return ""
}

// Int64ToHex int64转十六进制
func Int64ToHex(num int64) string {
	return strconv.FormatInt(num, 16)
}

// HexToString 十六进制转字符串
func HexToString(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}

// Round2 保留n位小数
func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

// IntToFloat64 .
func IntToFloat64(n int) float64 {
	d, _ := strconv.ParseFloat(strconv.Itoa(n), 64)
	return d
}

// StringToInt64 .
func StringToInt64(n string) int64 {
	d, _ := strconv.ParseInt(n, 10, 64)
	return d
}

// StringToUint .
func StringToUint(n string) uint {
	d, _ := strconv.ParseInt(n, 10, 64)
	return uint(d)
}

// UintToString .
func UintToString(n uint) string {
	return strconv.Itoa(int(n))
}

// StringToFloat64 .
func StringToFloat64(n string) float64 {
	d, _ := strconv.ParseFloat(n, 64)
	return d
}

// FormatIntPrefixNum .
func FormatIntPrefixNum(num int, prefix string, n int) string {
	return fmt.Sprintf("%"+prefix+strconv.Itoa(n)+"d", num)
}

// NegativeFloat64 .
func NegativeFloat64(n float64) float64 {
	return 0 - n
}

// FormatFloat64 value保留2位小数
func FormatFloat64(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// FormatFloat 无四舍五入
func FormatFloat(data float64, n int) float64 {
	temp := fmt.Sprintf("%v", data)
	list := strings.Split(temp, ".")
	if len(list) < 2 {
		return data
	}
	length := len(list[1])
	if length <= n {
		return data
	}
	list[1] = string([]byte(list[1])[:n])
	theString := strings.Join(list, ".")
	theFloat64, _ := strconv.ParseFloat(theString, 64)
	return theFloat64
}

// FloatToStr float金额转string
func FloatToStr(money float64) string {
	return strconv.FormatFloat(money, 'f', 2, 64)
}

// EarthDistance 经纬度计算两点之间的距离(米)
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
