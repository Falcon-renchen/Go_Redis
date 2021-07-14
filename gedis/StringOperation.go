package gedis

import (
	"context"
)

//专门处理string类型的操作
type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (this *StringOperation) Set() {

}

//如果有错，交给此函数来处理
func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

//获取多值
func (this *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, keys...).Result())
}
