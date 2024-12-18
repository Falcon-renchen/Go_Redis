package gedis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

const (
	Serilizer_JSON = "json"
	Serilizer_GOB  = "gob"
)

type DBGettFuc func() string

type SimpleCache struct {
	Operation *StringOperation // 操作类
	Expire    time.Duration    //过期时间
	DBGetter  DBGettFuc        //一旦缓存没有  DB的获取方法
	Serilizer string           // 序列化方式
	Policy    CachePolicy
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serilizer string, policy CachePolicy) *SimpleCache {
	policy.SetOperation(operation)
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer, Policy: policy}
}

//设置缓存
func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

func (this *SimpleCache) GetCacheForObject(key string, obj interface{}) interface{} {
	ret := this.GetCache(key)
	if ret == nil {
		return nil
	}
	if this.Serilizer == Serilizer_JSON {
		err := json.Unmarshal([]byte(ret.(string)), obj)
		if err != nil {
			return err
		} else if this.Serilizer == Serilizer_GOB {
			var buf = &bytes.Buffer{}
			buf.WriteString(ret.(string))
			dec := gob.NewDecoder(buf)
			if dec.Decode(obj) != nil {
				return nil
			}
		}
	}
	return nil
}

//gin 做测试
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	if this.Policy != nil {
		this.Policy.Before(key)
	}
	obj := this.DBGetter()
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
	} else if this.Serilizer == Serilizer_GOB {
		f := func() string {
			var buf = &bytes.Buffer{}
			enc := gob.NewEncoder(buf)

			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f)
		//this.SetCache(key,ret)
	}
	if ret.(string) == "" && this.Policy != nil {
		this.Policy.IfNil(key, "")
	} else {
		this.SetCache(key, ret)
	}
	return
}
