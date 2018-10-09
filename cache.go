package opentaobao

import "time"

// GetCache 获取缓存委托
type GetCacheFunc func(cacheKey string) []byte

// SetCache 设置缓存委托
type SetCacheFunc func(key string, value []byte, expiration time.Duration) string
