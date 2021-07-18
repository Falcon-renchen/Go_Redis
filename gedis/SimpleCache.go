package gedis

import (
	"encoding/json"
	"time"
)

const (
	Serilizer_JSON = "json"
)

type DBGettFuc func() string

type SimpleCache struct {
	Operation *StringOperation // 操作类
	Expire    time.Duration    //过期时间
	DBGetter  DBGettFuc        //一旦缓存没有  DB的获取方法
	Serilizer string           // 序列化方式
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serilizer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer}
}

//设置缓存
func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

//gin 做测试
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	if this.Serilizer == Serilizer_JSON {
		f := func() string {
			obj := this.DBGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f)
		this.SetCache(key, ret)
	}

	return
}
