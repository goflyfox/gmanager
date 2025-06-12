// Package cache
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cache

import (
	"context"
	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
)

// cache 缓存驱动
var cache gtoken.Cache

// Instance 缓存实例
func Instance() gtoken.Cache {
	if cache == nil {
		panic("cache nil")
	}
	return cache
}

// InitCache 初始化缓存
func InitCache(ctx context.Context) {
	if cache == nil {
		mode, _ := g.Cfg().Get(ctx, "cache.mode", gtoken.CacheModeFile)
		preKey, _ := g.Cfg().Get(ctx, "cache.preKey", "gmanager:")
		timeOut, _ := g.Cfg().Get(ctx, "cache.timeOut", 0)
		cache = gtoken.NewDefaultCache(mode.Int8(), preKey.String(), timeOut.Int64())
	}
}
