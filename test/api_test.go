package main

import (
	"fmt"
	"github.com/gogf/gf/g/util/gconv"
	"testing"
)

func TestRun(t *testing.T) {
	foreach("1", "2")
	var paramsStr []string
	paramsStr = append(paramsStr, "3")
	paramsStr = append(paramsStr, "4")
	foreach(paramsStr)

	var params []interface{}
	params = append(params, "5")
	params = append(params, "6")
	foreach(params)
	foreach(params...)

}

func foreach(args ...interface{}) {
	for _, a := range args {
		fmt.Println("#" + gconv.String(a))
	}
}
