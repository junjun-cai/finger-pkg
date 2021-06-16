//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/05/13 14:41
//@File:jlog.go
//@Pack:jlog
//@Proj:finger-pkg
//@Ides:GoLand
//@Desc:
//*********************************************************************************
package jlog

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	dailyRolling bool = true
)

const dateFormat = "20060102"

//Auth:2021/05/13 14:45:18 周四 cole-cai
//Desc:
type _Jlog struct {
	dir      string
	fileName string
	_suffix  int
	_date    *time.Time
	mu       *sync.RWMutex
	logfiel  *os.File
	lg       *log.Logger
}

//Auth:2021/05/13 14:53:14 周四 cole-cai
//Desc:


//Auth:2021/05/13 14:46:14 周四 cole-cai
//Desc:
func (j *_Jlog) isNeedRename() bool {
	if dailyRolling {
		t, _ := time.Parse(dateFormat, time.Now().Format(dateFormat))
		if t.After(*j._date) {
			return true
		}
	}


	return false
}
