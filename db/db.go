package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/junjun-cai/finger-pkg/conf"
	"github.com/junjun-cai/finger-pkg/obj"
	"gopkg.in/mgo.v2"
	"strings"
	"time"
)

//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/3/30 12:11
//@File:db.go
//@Pack:base
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

//Auth:2021-03-30 22:25:14 周二 cole-cai
//Desc:连接redis
func ConnectRedis(sec string) (*redis.Pool, error) {
	d := &obj.DialInfo{}
	e := conf.LoadSection(sec, d)
	if e != nil {
		return nil, e
	}
	return &redis.Pool{
		MaxIdle:     d.MaxIdle,
		MaxActive:   d.MaxActive,
		IdleTimeout: time.Duration(d.MaxTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial("tcp", d.Host)
			if e != nil {
				return nil, e
			}
			if strings.Compare("", d.Password) != 0 {
				if _, e = c.Do("AUTH", d.Password); e != nil {
					c.Close()
					return nil, e
				}
			}
			return c, e
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, e := c.Do("PING")
			return e
		},
	}, nil
}

//Auth:2021-03-30 22:29:32 周二 cole-cai
//Desc:连接mongo
func ConnectMongo(sec string) (*mgo.Session, error) {
	d := &obj.DialInfo{}
	e := conf.LoadSection(sec, d)
	if e != nil {
		return nil, e
	}
	c, e := mgo.Dial(d.Host)
	if e != nil {
		return nil, e
	}
	c.SetMode(mgo.Monotonic, true)
	return c, nil
}

//Auth:2021-03-30 22:32:05 周二 cole-cai
//Desc:连接mySQL
func ConnectMySQL(sec string) (*gorm.DB, error) {
	d := &obj.DialInfo{}
	e := conf.LoadSection(sec, d)
	if e != nil {
		return nil, e
	}

	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.DataBase)
	g, e := gorm.Open("mysql", s)
	if e != nil {
		return nil, e
	}
	g.SingularTable(true)
	g.DB().SetMaxIdleConns(d.MaxIdle)
	g.DB().SetMaxOpenConns(d.MaxActive)
	return g, nil
}
