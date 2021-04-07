package utils

import (
	"fmt"
	"github.com/junjun-cai/finger-pkg/define"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
	"math/rand"
	"regexp"
	"time"
)

//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/3/30 12:31
//@File:utils.go
//@Pack:utils
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

//Auth:2021-03-30 15:39:27 周二 cole-cai
//Desc
func BKDRHash(str string) int {
	seed := 13131 // 31 131 1313 13131 131313 etc..
	hash := 0
	for _, ch := range str {
		hash = hash*seed + int(ch)
	}
	return hash & 0x7FFFFFFF
}

//Auth:2021-03-30 15:39:45 周二 cole-cai
//Desc: eg:20060102
func TodayDateStr() string {
	return time.Now().Format(define.DateFormat)
}

//Auth:2021-03-30 15:40:08 周二 cole-cai
//Desc:判断连个时间戳是否同一天
func IsSameDay(tm1, tm2 int64) bool {
	t1 := time.Unix(tm1, 0)
	t2 := time.Unix(tm2, 0)
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

//Auth:2021-03-30 15:40:36 周二 cole-cai
//Desc:判断给定时间戳是否今天
func IsToday(tm int64) bool {
	t1 := time.Now()
	t2 := time.Unix(tm, 0)
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

//Auth:2021-03-30 15:41:28 周二 cole-cai
//Desc:判断给定时间是否当前周
func IsCurWeek(tm int64) bool {
	t1, w1 := time.Now().ISOWeek()
	t2, w2 := time.Unix(tm, 0).ISOWeek()
	return t1 == t2 && w1 == w2
}

//Auth:2021-03-30 15:42:01 周二 cole-cai
//Desc:获取当天剩余的秒数
func TodayLeftSecond() int64 {
	t1 := time.Now()
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 0, t1.Location())
	return t2.Unix() - t1.Unix()
}

//Auth:2021-03-30 15:42:32 周二 cole-cai
//Desc:根据message name 生成对应obj
func NewProtoMessage(name string) (proto.Message, error) {
	m, e := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(name))
	if e != nil {
		return nil, nil
	}
	return m.New().Interface(), nil
}

//Auth:2021-03-30 15:43:01 周二 cole-cai
//Desc:获取message.fullname
func GetMessageName(m proto.Message) string {
	return string(protoimpl.X.MessageDescriptorOf(m).FullName())
}

//Auth:2021/04/01 23:57:16 周四 cole-cai
//Desc:生成验证码
func GenValidateCode(width int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < width; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//Auth:2021/04/02 20:02:05 周五 cole-cai
//Desc:
func FormatValidateCodeContent(code string, expire int64) string {
	return fmt.Sprintf("[FingerGame]This is your verification code:%s,it will expire in %d seconds.", code, expire)
}

//Auth:2021/04/02 20:13:17 周五 cole-cai
//Desc:校验邮箱地址的合法性
func ValidateEmail(email string) bool {
	pattern := "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//Auth:2021/04/02 20:39:04 周五 cole-cai
//Desc:校验手机号的合法性
func ValidatePhone(phone string) bool {
	pattern := "^1(3\\d{2}|4[14-9]\\d|5([0-35689]\\d|7[1-79])|66\\d|7[2-35-8]\\d|8\\d{2}|9[13589]\\d)\\d{7}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

//Auth:2021/04/04 23:06:22 周日 cole-cai
//Desc:
func RetryTimes(name string, times int, duration time.Duration, fn func() error) error {
	for i := 0; i < times; i++ {
		if err := fn(); err == nil {
			return nil
		}
		time.Sleep(duration)
	}
	return fmt.Errorf("%s retry %d times,still failed", name, times)
}
