package define

//*********************************************************************************
//@Auth:cole-cai
//@Date:2021/3/30 15:23
//@File:const.go
//@Pack:define
//@Proj:base
//@Ides:GoLand
//@Desc:
//*********************************************************************************
//Auth:2021/03/30 22:05:29 周二 cole-cai
//Desc:
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
	DateFormat   string = "20060102"
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)
