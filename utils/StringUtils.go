package utils

import "fmt"
import "crypto/md5"
import "math/rand"
import "time"
import "strconv"
import (
	"strings"
	"regexp"
)

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

//拼接sql where in 的占位符
func GetWhereInSqlByStrId(strId string) (string, []string){
	idArr := strings.Split(strId, ",")
	t := make([]string, len(idArr))
	for i, _ := range idArr {
		t[i] = " ? "
	}
	return strings.Join(t, ","), idArr
}


//切割由数字组成的字符串到切片
func StringsSplitToSliceInt(s string, sep string) []int64 {
	if s == "" || sep == "" {
		return []int64{}
	}

	p := fmt.Sprintf(`^(\d+[%s]?)+\d|\d$`, sep)
	match,_:=regexp.MatchString(p,s)
	if match == false {
		return []int64{}
	}

	strArr := strings.Split(s, sep)
	intArr := make([]int64, len(strArr))

	for k,v := range strArr{
		intArr[k], _ = strconv.ParseInt(v, 10, 64)
	}

	return intArr
}

//过滤字符串
func TrimString(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	for i, ch := range r {
		switch {
		case ch == '\'':
			r[i] = 0
		case ch == '\r':
			r[i] = 0
		case ch == '\n':
			r[i] = 0
		case ch == '\t':
			r[i] = 0
		case ch == '`':
			r[i] = 0
		case ch == '"':
			r[i] = 0
		}
	}
	return string(r)
}

//日期转时间戳
func GetTimestamp(d string) int64 {
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", d, loc)
	if err == nil {
		return the_time.Unix()
	}
	return 0
}
