package utils

import (
	"github.com/junjun-cai/finger-pkg/define"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
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
