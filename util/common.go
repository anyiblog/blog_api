//工具公用函数库

package util

import (
	md52 "crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//随机字符串种子
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

const letterNumbers = "1234567890"

// RandStr 生成随机字符串
func RandStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandNumber 生成随机数据
func RandNumber(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterNumbers[rand.Intn(len(letterNumbers))]
	}
	return string(b)
}

// GenerateOrderNo 生成订单号，日期年月日+随机数字（长度15位）
func GenerateOrderNo() string {
	month := strconv.Itoa(int(time.Now().Month()))
	day := strconv.Itoa(time.Now().Day())
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	return fmt.Sprintf("%s%s%s%s", strconv.Itoa(time.Now().Year()), month, day, RandNumber(7))
}

// GenerateToken 生成Token
func GenerateToken() string {
	return RandStr(28) + "+" + RandStr(8) + "+" + RandStr(26)
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// Md5 生成MD5
func Md5(str string) string {
	data := []byte(str)
	md5Ctx := md52.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	strMd5 := fmt.Sprintf("%x", cipherStr)
	return strMd5
}

// StrJsonToMap 字符串JSON转map
func StrJsonToMap(strJson string) map[string]string {
	var strMap map[string]string
	if err := json.Unmarshal([]byte(strJson), &strMap); err != nil {
		fmt.Println(err)
	}
	return strMap
}

// MinToSm 返回多少天后的秒数，Redis缓存会用
func MinToSm(day string) int {
	now := time.Now()
	ad, _ := time.ParseDuration("24h")
	fmt.Println(now.Add(ad * 7).Unix())
	i := now.Add(ad*7).Unix() - now.Unix()
	return int(i)
}

// DayToSm 返回多少天后的秒数，Redis缓存会用
func DayToSm(day string) int {
	now := time.Now()
	ad, _ := time.ParseDuration("24h")
	fmt.Println(now.Add(ad * 7).Unix())
	i := now.Add(ad*7).Unix() - now.Unix()
	return int(i)
}

// StructToJson 结构体转json
func StructToJson(structData interface{}) string {
	jsons, errs := json.Marshal(structData) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return string(jsons)
}

// CheckPhone 验证手机号
func CheckPhone(phone string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// SixRandNum 随机6位数字码
func SixRandNum() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

// TimeFormat 时间类型 转 字符串类型
func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// StrTimeFormat 时间字符串 转 Time类型
func StrTimeFormat(strTime, timeLayout string) time.Time {
	theTime, err := time.ParseInLocation(timeLayout, strTime, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	return theTime
}

// NowTime 当前时间 数据库插入使用
func NowTime() time.Time {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Now().Format(timeLayout)
	loc, _ := time.LoadLocation("Local")                         //时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //字符串转换Time类型
	return theTime
}

// Fen2Yuan 分转元,
func Fen2Yuan(price string) float64 {
	priceInt, _ := strconv.Atoi(price)
	d := decimal.New(1, 2) //分除以100得到元
	result, _ := decimal.NewFromInt(int64(priceInt)).DivRound(d, 2).Float64()
	return result
}

// Yuan2Fen 元转分,乘以100后，保留整数部分
func Yuan2Fen(price float64) int64 {
	d := decimal.New(1, 2)  //分转元乘以100
	d1 := decimal.New(1, 0) //乘完之后，保留2为小数，需要这么一个中间参数
	//如下是满足，当乘以100后，仍然有小数位，取四舍五入法后，再取整数部分
	dff := decimal.NewFromFloat(price).Mul(d).DivRound(d1, 0).IntPart()
	return dff
}
