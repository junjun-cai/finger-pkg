package define

//*********************************************************************************
//@Auth:蔡君君
//@Date:2021/3/30 15:23
//@File:const.go
//@Pack:define
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************

const (
	MinuteSecond        = 60
	HourSecond   int64  = MinuteSecond * 60
	DaySecond    int64  = HourSecond * 24
	WeekSecond   int64  = DaySecond * 7
	HourMinute   int64  = 60
	DayMinute    int64  = HourMinute * 24
	WeekMinute   int64  = DayMinute * 7
	MoneyK              = 1000
	MoneyM              = 1000 * MoneyK
	MoneyB              = 1000 * MoneyM
	MoneyT              = 1000 * MoneyB
	DateFormat    string = "20060102"
)
