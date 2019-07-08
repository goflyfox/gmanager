package gtoken

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/os/gcache"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils/resp"
)

// setCache 设置缓存
func (m *GfToken) setCache(cacheKey string, userCache g.Map) resp.Resp {
	switch m.CacheMode {
	case CacheModeCache:
		gcache.Set(cacheKey, userCache, m.Timeout)
	case CacheModeRedis:
		cacheValueJson, err1 := gjson.Encode(userCache)
		if err1 != nil {
			glog.Error("[GToken]cache json encode error", err1)
			return resp.Error("cache json encode error")
		}
		_, err := g.Redis().Do("SETEX", cacheKey, m.Timeout, cacheValueJson)
		if err != nil {
			glog.Error("[GToken]cache set error", err)
			return resp.Error("cache set error")
		}
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ(userCache)
}

// getCache 获取缓存
func (m *GfToken) getCache(cacheKey string) resp.Resp {
	var userCache g.Map
	switch m.CacheMode {
	case CacheModeCache:
		userCacheValue := gcache.Get(cacheKey)
		if userCacheValue == nil {
			return resp.Unauthorized("login timeout or not login", "")
		}
		userCache = gconv.Map(userCacheValue)
	case CacheModeRedis:
		userCacheJson, err := g.Redis().Do("GET", cacheKey)
		if err != nil {
			glog.Error("[GToken]cache get error", err)
			return resp.Error("cache get error")
		}
		if userCacheJson == nil {
			return resp.Unauthorized("login timeout or not login", "")
		}

		err = gjson.DecodeTo(userCacheJson, &userCache)
		if err != nil {
			glog.Error("[GToken]cache get json error", err)
			return resp.Error("cache get json error")
		}
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ(userCache)
}

// removeCache 删除缓存
func (m *GfToken) removeCache(cacheKey string) resp.Resp {
	switch m.CacheMode {
	case CacheModeCache:
		gcache.Remove(cacheKey)
	case CacheModeRedis:
		var err error
		_, err = g.Redis().Do("DEL", cacheKey)
		if err != nil {
			glog.Error("[GToken]cache remove error", err)
			return resp.Error("cache remove error")
		}
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ("")
}
