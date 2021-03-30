package obj

import "fmt"

//*********************************************************************************
//@Auth:蔡君君
//@Date:2021/3/30 15:11
//@File:obj.go
//@Pack:obj
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

type UserBase struct {
	Mid      int64 `bson:"mid" json:"mid" redis:"mid" gorm:"column:mid"`
	Channel  int32 `bson:"channel" json:"channel" redis:"channel" gorm:"column:channel"`
	AppId    int32 `bson:"app_id" json:"app_id" redis:"app_id" gorm:"column:app_id"`
	Platform int32 `bson:"platform" json:"platform" redis:"platform" gorm:"column:platform"`
}

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
	PhoneNumber  string           `bson:"phone_number" json:"phone_number" redis:"phone_number" gorm:"column:phone_number"`
	Version      string           `bson:"version" json:"version" redis:"version" gorm:"column:version"`
	BackPack     string           `bson:"back_pack" json:"back_pack" redis:"back_pack" gorm:"column:back_pack"`
	GameInfo     string           `bson:"game_info" json:"game_info" redis:"game_info" gorm:"column:game_info"`
}

func (u User) TableName() string {
	return fmt.Sprintf("user_%d", u.Mid%10)
}

type Register struct {
	UserBase `bson:",inline"`
	Uuid     string `bson:"uuid" json:"uuid" redis:"uuid" gorm:"column:uuid"`
	UuidHash int    `bson:"uuid_hash" json:"uuid_hash" redis:"uuid_hash" gorm:"column:uuid_hash"`
}

func (r Register) TableName() string {
	return fmt.Sprintf("register_%d", r.UuidHash%10)
}

type Asset struct {
	UserBase `bson:",inline"`
	Money    int64 `bson:"money" json:"money" redis:"money" gorm:"column:money"`
}

func (a Asset) TableName() string {
	return fmt.Sprintf("asset_%d", a.Mid%10)
}

type Friends struct {
	Mid     int64   `bson:"mid" json:"mid" redis:"mid" gorm:"column:mid"`
	Friends []int64 `bson:"friends" json:"friends" redis:"friends" gorm:"column:friends"`
}


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