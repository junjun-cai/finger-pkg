package base

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//*********************************************************************************
//@Auth:蔡君君
//@Date:2021/3/30 12:25
//@File:conf.go
//@Pack:base
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

var i *ini.File

func LoadIniConfig(file string) error {
	var e error
	i, e = ini.Load(file)
	if e != nil {
		return e
	}
	return nil
}

func LoadSection(sec string, conf interface{}) error {
	if i == nil {
		return fmt.Errorf("cant read section,load config must firstly")
	}
	if s, e := i.GetSection(sec); e != nil {
		return e
	} else {
		return s.MapTo(conf)
	}
}
