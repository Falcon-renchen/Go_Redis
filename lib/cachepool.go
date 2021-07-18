package lib

import (
	"Go_Redis/gedis"
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init() {
	NewsCachePool = &sync.Pool{
		New: func() interface{} {
			return gedis.NewSimpleCache(gedis.NewStringOperation(), time.Second*15,
				gedis.Serilizer_JSON) //指定序列化，是JSON
		},
	}
}
func NewsCache() *gedis.SimpleCache {
	return NewsCachePool.Get().(*gedis.SimpleCache)
}

func ReleaseNewsCache(cache *gedis.SimpleCache) {
	NewsCachePool.Put(cache)
}
