package main

import (
	"fmt"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

// 判断是否工作时间
func isWorkingPeriod(t *gtime.Time) bool {
	_, weekNumber := t.ISOWeek()
	if weekNumber == 5 || weekNumber == 6 {
		return false
	}
	if t.Hour() < 9 && t.Hour() > 18 {
		return false
	}
	return true
}

func main() {
	startTime := gtime.NewFromStr("2018-06-01")
	match, _ := gregex.MatchString(`@(\d+)`, genv.Get("GIT_COMMITTER_DATE"))
	if len(match) > 1 {
		commitTime := gtime.NewFromTimeStamp(gconv.Int64(match[1]))
		if commitTime.After(startTime) {
			if isWorkingPeriod(commitTime) {
				fmt.Print(commitTime.FormatTo(fmt.Sprintf(`Y-m-d %d:i:s`, grand.N(18, 23))).RFC822())
			}
		}
	}
}
