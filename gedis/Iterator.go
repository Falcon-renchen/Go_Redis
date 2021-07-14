package gedis

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

//是否有值
func (this *Iterator) HasNext() bool {
	if this.data == nil || len(this.data) == 0 {
		return false
	}
	return this.index < len(this.data)
}

//取值
func (this *Iterator) Next() (ret interface{}) {
	ret = this.data[this.index]
	this.index = this.index + 1
	return
}
