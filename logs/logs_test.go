//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/05/13 14:56
//@File:logs_test.go
//@Pack:logs
//@Proj:finger-pkg
//@Ides:GoLand
//@Desc:
//*********************************************************************************
package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestLogs(t *testing.T) {
	tm := time.NewTicker(1 * time.Second)
	for{
		select {
		case  <- tm.C:
			fmt.Println("xx")
		}
	}
}
