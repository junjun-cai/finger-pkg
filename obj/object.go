package obj

import "fmt"

//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/3/30 15:11
//@File:obj.go
//@Pack:obj
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

//Auth:2021/03/30 22:05:37 周二 cole-cai
//Desc:
type UserBase struct {
	Mid      int64 `bson:"mid" json:"mid" redis:"mid" gorm:"column:mid"`
	Channel  int32 `bson:"channel" json:"channel" redis:"channel" gorm:"column:channel"`
	AppId    int32 `bson:"app_id" json:"app_id" redis:"app_id" gorm:"column:app_id"`
	Platform int32 `bson:"platform" json:"platform" redis:"platform" gorm:"column:platform"`
}

//Auth:2021/03/30 22:06:55 周二 cole-cai
//Desc:
type User struct {
	UserBase     `bson:",inline"` //mongodb 中组合要用inline tag,否则会将组合当做一个嵌套的结构体使用，而不会展开
	Sex          int32            `bson:"sex" json:"sex" redis:"sex" gorm:"column:sex"`
	LoginCnt     int32            `bson:"login_cnt" json:"login_cnt" redis:"login_cnt" gorm:"column:login_cnt"`
	RegisterTime int64            `bson:"register_time" json:"register_time" redis:"register_time" gorm:"column:register_time"`
	LastLogin    int64            `bson:"last_login" json:"last_login" redis:"last_login" gorm:"column:last_login"`
	Pay          float64          `bson:"pay" json:"pay" redis:"pay" gorm:"column:pay"`
	Name         string           `bson:"name" json:"name" redis:"name" gorm:"column:name"`
	Describe     string           `bson:"describe" json:"describe" redis:"describe" gorm:"column:describe"`
	Icon         string           `bson:"icon" json:"icon" redis:"icon" gorm:"column:icon"`
	Uuid         string           `bson:"uuid" json:"uuid" redis:"uuid" gorm:"column:uuid"`
	Password     string           `bson:"password" json:"password" redis:"password" gorm:"column:password"`
	Phone        string           `bson:"phone" json:"phone" redis:"phone" gorm:"column:phone"`
	Email        string           `bson:"email" json:"email" redis:"email" gorm:"column:email"`
	Version      string           `bson:"version" json:"version" redis:"version" gorm:"column:version"`
	BackPack     string           `bson:"back_pack" json:"back_pack" redis:"back_pack" gorm:"column:back_pack"`
	GameInfo     string           `bson:"game_info" json:"game_info" redis:"game_info" gorm:"column:game_info"`
}

//Auth:2021/03/30 22:12:08 周二 cole-cai
//Desc:
func (u User) TableName() string {
	return fmt.Sprintf("user_%d", u.Mid%10)
}

//Auth:2021/03/30 22:12:33 周二 cole-cai
//Desc:
type Register struct {
	UserBase `bson:",inline"`
	Uuid     string `bson:"uuid" json:"uuid" redis:"uuid" gorm:"column:uuid"`
	UuidHash int    `bson:"uuid_hash" json:"uuid_hash" redis:"uuid_hash" gorm:"column:uuid_hash"`
}

//Auth:2021/03/30 22:13:35 周二 cole-cai
//Desc:
func (r Register) TableName() string {
	return fmt.Sprintf("register_%d", r.UuidHash%10)
}

//Auth:2021/03/30 22:13:48 周二 cole-cai
//Desc:
type Asset struct {
	UserBase `bson:",inline"`
	Money    int64 `bson:"money" json:"money" redis:"money" gorm:"column:money"`
}

//Auth:2021/03/30 22:14:01 周二 cole-cai
//Desc:
func (a Asset) TableName() string {
	return fmt.Sprintf("asset_%d", a.Mid%10)
}

//Auth:2021/03/30 22:15:12 周二 cole-cai
//Desc:
type Friends struct {
	Mid     int64   `bson:"mid" json:"mid" redis:"mid" gorm:"column:mid"`
	Friends []int64 `bson:"friends" json:"friends" redis:"friends" gorm:"column:friends"`
}

//Auth:2021/03/30 22:17:39 周二 cole-cai
//Desc:
func (f Friends) TableName() string {
	return fmt.Sprintf("friend_%d", f.Mid%10)
}

//Auth:2021/03/30 22:19:25 周二 cole-cai
//Desc:
type DialInfo struct {
	User       string
	Host       string
	Password   string
	DataBase   string
	MaxIdle    int
	MaxActive  int
	PoolLimit  int
	MaxTimeout int64
}

//Auth:2021/03/30 22:28:30 周二 cole-cai
//Desc:
type App struct {
	Name    string
	Host    string
	RunMode string
	Crontab bool
}

//Auth:2021/03/30 22:29:01 周二 cole-cai
//Desc:
type Log struct {
	SavePath string
	SaveName string
	FileSize int
	OutPut   int
	Level    int
}

//Auth:2021/04/01 20:46:46 周四 cole-cai
//Desc:
type GoldLog struct {
	Payer    int64  `json:"payer"`
	PayerOld int64  `json:"payer_old"`
	PayerNew int64  `json:"payer_new"`
	Payee    int64  `json:"payee"`
	PayeeOld int64  `json:"payee_old"`
	PayeeNew int64  `json:"payee_new"`
	Amount   int64  `json:"amount"`
	Time     int64  `json:"time"`
	Act      int32  `json:"act"`
	Info     string `json:"info"`
	LogId    string `json:"id"`
}

//Auth:2021/04/01 23:28:02 周四 cole-cai
//Desc:
type WebLog struct {
	Mid      int64  `json:"mid"`
	Appid    int32  `json:"appid"`
	Channel  int32  `json:"channel"`
	Platform int32  `json:"platform"`
	Version  string `json:"version"`
	LogId    string `json:"log_id"`
	Info     string `json:"msg,omitempty"`
}
