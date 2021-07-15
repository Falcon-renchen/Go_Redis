package gedis

import "time"

type DBGettFuc func() string

type SimpleCache struct {
	Operation *StringOperation
	Expire    time.Duration
	DBGetter  DBGettFuc
}

func NewSimpleCache(operation *StringOperation, expire time.Duration) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire}
}

//设置缓存
func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

//gin 做测试
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	ret = this.Operation.Get(key).Unwrap_Or_Else(this.DBGetter)
	this.SetCache(key, ret)
	return
}
