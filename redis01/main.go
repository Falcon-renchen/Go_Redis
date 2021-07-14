package main

import (
	"Go_Redis/gedis"
	"fmt"
)

func main() {
	//ctx := context.Background()
	//
	//ret := gedis.Redis().Get(ctx,"name")
	//v, err := ret.Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(v)

	//fmt.Println(gedis.NewStringOperation().Get("name").Unwrap_Or("default value"))   //执行结果： shenyi

	iter := gedis.
		NewStringOperation().             // string类型的处理类
		MGet("name", "age", "abc").Iter() //变成自己的迭代器

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}

}
