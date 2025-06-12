package logic_test

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func TestDate(t *testing.T) {
	var dates []string
	startDate := "2025-05-29"
	endDate := "2025-06-03"
	startTime := gtime.NewFromStrFormat(startDate, "Y-m-d")
	endTime := gtime.NewFromStrFormat(endDate, "Y-m-d")
	for !startTime.After(endTime) {
		dates = append(dates, startTime.Format("Y-m-d"))
		startTime = startTime.AddDate(0, 0, 1)
	}
	fmt.Println(dates)
}
