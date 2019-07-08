package main

import (
	"github.com/gogf/gf/g/net/ghttp"
	"testing"
)

const (
	TestURL string = "http://127.0.0.1:80"
)

func TestLogTest(t *testing.T) {
	if r, e := ghttp.Get(TestURL + "/tmp/test?name=john&age=18"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestLogPost(t *testing.T) {
	if r, e := ghttp.Post(TestURL+"/tmp/access", "name=john&age=18"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestLogGet(t *testing.T) {
	if r, e := ghttp.Get(TestURL + "/tmp/access?name=john&age=18"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestLogCache(t *testing.T) {
	if r, e := ghttp.Get(TestURL + "/tmp/cache"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestLogMysql(t *testing.T) {
	if r, e := ghttp.Get(TestURL + "/tmp/mysql"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}
