package main

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestUserGet(t *testing.T) {
	if r, e := g.Client().Get(TestURL + "/system/user/get?id=1"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestUserInsert(t *testing.T) {
	if r, e := g.Client().Get(TestURL + "/system/user/access?name=john&age=18"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestUserUpdate(t *testing.T) {
	if r, e := g.Client().Get(TestURL + "/system/user/cache"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestUserDelete(t *testing.T) {
	if r, e := g.Client().Post(TestURL+"/system/user/access", "name=john&age=18"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestUserList(t *testing.T) {
	if r, e := g.Client().Get(TestURL + "/system/user/mysql"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}

func TestUserPage(t *testing.T) {
	if r, e := g.Client().Get(TestURL + "/system/user/mysql"); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}
