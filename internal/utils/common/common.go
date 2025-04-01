package common

import (
	"fmt"
	"game/internal/utils/model"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

func GetTimeMicro() int64 {
	return time.Now().UnixMicro()
}
func GetTimeNano() int64 {
	return time.Now().UnixNano()
}
func GetTimeMs() int64 {
	return time.Now().UnixMilli()
}
func GetTime() int64 {
	return time.Now().Unix()
}
func GetTimeUTC() time.Time {
	return time.Now().UTC()
}
// GetTimeUTCN 获取 UTC+n 的时间
func GetTimeUTCN(n int) time.Time {
	utcTime := time.Now().UTC()
	utcNZone := time.FixedZone(fmt.Sprintf("UTC+%d", n), n*60*60)
	return utcTime.In(utcNZone)
}
// GetTomorrowMidnightUTC 获取明天凌晨的UTC时间
func GetTomorrowMidnightUTC() time.Time {
    now := time.Now().UTC()
    tomorrow := now.AddDate(0, 0, 1)
    tomorrowMidnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.UTC)
    return tomorrowMidnight
}


// TimeToUTCN 将时间转换为 UTC+n 的时间
func TimeToUTCN(t time.Time, n int) time.Time {
	utcNZone := time.FixedZone(fmt.Sprintf("UTC+%d", n), n*60*60)
	return t.In(utcNZone)
}

func UTCToISOWeek(t time.Time) string {
	year, week := t.ISOWeek()
    return fmt.Sprintf("%d%d", year, week) 
}

func UTCToISODay(t time.Time) string {
	return t.Format("20060102")
}

// 获取当天的零点
func GetTodayZeroTime() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func GetNumIsEven[T model.Number](data T) bool {
	return int64(data)&1 == 0
}

// DoubleToString 小数类型转 string类型
func DoubleToString[T model.Float](data T) string {
	return strconv.FormatFloat(float64(data), 'f', 2, 64)
}

// NumberToString 数字类型转 string类型
func NumberToString[T model.Number](data T) string {
	return strconv.FormatInt(int64(data), 10)
}

// StringToNumber string类型转 数字类型
func StringToNumber[T model.Number](data string) T {
	s, _ := strconv.ParseInt(data, 10, 64)
	return T(s)
}

// Marshal 结构体转json格式字符串
func Marshal(data any) (s string) {
	s, _ = json.MarshalToString(data)
	return s
}

// Unmarshal json格式字符串转结构体
func Unmarshal(s string, data any) error {
	return json.UnmarshalFromString(s, data)
}

// GetOnlyId 获取唯一id
func GetOnlyId() string {
	id := uuid.New().String()
	return strings.ReplaceAll(id, "-", "")
}

// GetKey 拼接key
func GetKey(key string, field ...string) (s string) {
	data := strings.Builder{}
	data.WriteString(key)
	for _, v := range field {
		data.WriteString("_")
		data.WriteString(v)
	}
	return data.String()
}
func GetDayTimeFormat[T model.Number](day T) string {
	return time.Now().AddDate(0, 0, int(day)).Format("2006-01-02")
}

// GetTimeIsSame 判断时间是否是n天前/后
// day 偏移的天数
func GetTimeIsSame[T model.Number](times, day T) bool {
	now := time.Now().AddDate(0, 0, int(day)).Format("2006-01-02")
	sign := time.Unix(int64(times), 0).Format("2006-01-02")
	return now == sign
}

// GetExpDaySecond 获取n天前/后相差的秒数
// day 偏移的天数
func GetExpDaySecond[T model.Number](day T) T {
	timeStr := time.Now().Format("2006-01-02")
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if day > 0 {
		return T(t2.AddDate(0, 0, int(day)).Unix() - time.Now().Unix())
	} else {
		return T(time.Now().Unix() - t2.AddDate(0, 0, int(day)).Unix())
	}
}

// GetExpTime 获取偏移后的时间类型
func GetExpTime[T model.Number](bufferTime T) time.Time {
	return time.Now().Add(time.Duration(bufferTime))
}

// GetExpUnix 获取偏移后的时间戳
func GetExpUnix(day int) int64 {
	return time.Now().AddDate(0, 0, day).Unix()
}

// GetMonthDays 获取一个时间当月共有多少天
func GetMonthDays(t time.Time) int {
	t = GetToday(t)
	year, month, _ := t.Date()
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			return 30
		}
		return 31
	}

	if ((year%4 == 0) && (year%100 != 0)) || year%400 == 0 {
		return 29
	}

	return 28
}

// WeekDay 获取一个时间是星期几
//   - 1 ~ 7
func WeekDay(t time.Time) int {
	t = GetToday(t)
	week := int(t.Weekday())
	if week == 0 {
		week = 7
	}

	return week
}

// GetToday 获取一个时间的今天
func GetToday(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// GetSecond 获取共有多少秒
func GetSecond(d time.Duration) int {
	return int(d / time.Second)
}

// IsSameDay 两个时间是否是同一天
func IsSameDay(t1, t2 time.Time) bool {
	t1, t2 = GetToday(t1), GetToday(t2)
	return t1.Unix() == t2.Unix()
}

// IsSameHour 两个时间是否是同一小时
func IsSameHour(t1, t2 time.Time) bool {
	return t1.Hour() == t2.Hour() && t1.Day() == t2.Day() && t1.Month() == t2.Month() && t1.Year() == t2.Year()
}

/*
*
获取指定周一的时间
week:

	-1:上周
	 0:本周
	 1:下周

*
*/
func WeekIntervalTime(week int) (startTime string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	year, month, day := now.Date()
	thisWeek := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	startTime = thisWeek.AddDate(0, 0, offset+7*week).Format("2006-01-02")
	//endTime = thisWeek.AddDate(0, 0, offset+6+7*week).Format("2006-01-02")

	return startTime
}

// GetRandom 随机数区间
func GetRandom[T model.Number](min, max T) T {
	base := int64(min) + rand.Int63n(int64(max-min+1))
	return T(base)
}

// Decimal 精确浮点加减
// num 保留小数点后几位
func Decimal[T model.Float, A model.Number](data1, data2 T, types, num A) T {
	var value float64
	switch types {
	case model.Multiply:
		value, _ = decimal.NewFromFloat(float64(data1)).Mul(decimal.NewFromFloat(float64(data2))).Float64()
	case model.Add:
		value, _ = decimal.NewFromFloat(float64(data1)).Add(decimal.NewFromFloat(float64(data2))).Float64()
	case model.Sub:
		value, _ = decimal.NewFromFloat(float64(data1)).Sub(decimal.NewFromFloat(float64(data2))).Float64()
	case model.Divide:
		value, _ = decimal.NewFromFloat(float64(data1)).Div(decimal.NewFromFloat(float64(data2))).Float64()

	}
	return T(DecimalData(value, int32(num)))
}

// DecimalData 保留小数点后几位
func DecimalData(value float64, num int32) float64 {
	d, _ := decimal.NewFromFloat(value).RoundFloor(num).Float64()
	return d
}

// BinarySearch 两分法
func BinarySearch[T model.Number](data []T, target T) int {
	left, right := 0, len(data)
	for left <= right {
		mid := left + (right-left)/2
		if data[mid] > target {
			right = mid - 1
		} else if data[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || data[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

// 数组转 string字符串
func ArrayToString(array []any) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// 复制指针
func CopyPoint(m any) any {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface()
}

// 结构体转map
func StructToMapString(obj any) map[string]string {
	mapping := make(map[string]string)
	var valueOf = reflect.ValueOf(obj)
	if valueOf.Kind() == reflect.Pointer {
		valueOf = reflect.ValueOf(obj).Elem()
	}
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		jTag := valueOf.Type().Field(i).Tag.Get("json")
		index := strings.Index(jTag, ",")
		if index > 0 {
			jTag = jTag[:index]
		}
		if field.IsZero() {
			continue
		}
		mapping[jTag] = fmt.Sprint(field.Interface())
	}
	return mapping
}
