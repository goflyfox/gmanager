package util

import "github.com/gogf/gf/util/gconv"

func ToCols(obj interface{}) string {
	m := gconv.Map(obj)
	str := ""
	for _, v := range m {
		str += gconv.String(v) + ","
	}

	str = str[0 : len(str)-1]
	return str
}
