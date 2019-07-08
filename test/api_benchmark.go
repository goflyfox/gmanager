package main

import (
	"github.com/gogf/gf/g/net/ghttp"
	"testing"
)

const (
	BenchmarkURL = "http://127.0.0.1:80"
)

func BenchmarkLogGet(b *testing.B) {
	if r, e := ghttp.Get(BenchmarkURL + "/log/access?name=john&age=18"); e != nil {
		b.Log(e)
	} else {
		b.Log(string(r.ReadAll()))
		r.Close()
	}
}
