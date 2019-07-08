package utils

import "github.com/gogf/gf/g/os/gtime"

func GetNow() string {
	return gtime.Datetime()
}
