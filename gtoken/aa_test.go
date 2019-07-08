package gtoken

import (
	"github.com/gogf/gf/g/crypto/gdes"
	"github.com/gogf/gf/g/encoding/gbase64"
	"testing"
)

func TestLogTest(t *testing.T) {
	test := []byte("1123")
	key := []byte("d2eb867e67408065717eb274116ccc13")
	t.Log(gdes.DesECBEncrypt(key))
	eTest, _ := Encrypt(test, key)
	enStr := gbase64.Encode(string(eTest))
	t.Log(enStr)
	dTest, _ := Decrypt(eTest, key)
	t.Log(string(dTest))
}

const (
	ivDefValue = "I Love Go Frame!"
)
