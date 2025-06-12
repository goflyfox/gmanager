// Package captcha
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package captcha

import (
	"context"
	"image/color"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
)

// store 验证码存储方式
var store = base64Captcha.DefaultMemStore

// Generate 生成验证码
func Generate(ctx context.Context) (id string, base64 string, err error) {
	// 算数
	driver := &base64Captcha.DriverMath{
		Height:          42,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 255,
			G: 250,
			B: 250,
			A: 250,
		},
		Fonts: []string{"chromohv.ttf"},
	}

	c := base64Captcha.NewCaptcha(driver.ConvertFonts(), store)
	id, base64, _, err = c.Generate()
	if err != nil {
		g.Log().Errorf(ctx, "captcha.Generate err:%+v", err)
	}
	return
}

// Verify 验证输入的验证码是否正确
func Verify(id, answer string) bool {
	if id == "" || answer == "" {
		return false
	}
	return store.Verify(id, gstr.ToLower(answer), true)
}
