package main

import (
	"Go_Redis/lib"
	"github.com/gin-gonic/gin"
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

	//iter := gedis.
	//	NewStringOperation().             // string类型的处理类
	//	MGet("name", "age", "abc").Iter() //变成自己的迭代器
	//
	//for iter.HasNext() {
	//	fmt.Println(iter.Next())
	//}

	//fmt.Println(gedis.
	//	NewStringOperation().Get("name").Unwrap())
	//iter:=gedis.
	//	 NewStringOperation(). // string类型的处理类
	//	 MGet("name","age","abc").Iter() //变成自己的迭代器
	//
	//for iter.HasNext(){
	//	fmt.Println(iter.Next())
	//}

	//fmt.Println(gedis.
	//	NewStringOperation().
	//	Set("name","shenyi",gedis.WithExpire(time.Second*15)))

	////新闻缓存， 假设 我们认为他过期时间=15s
	//newsCache := gedis.NewSimpleCache(gedis.NewStringOperation(),time.Second*15)
	////新闻的缓存key:news123 news101
	//newsID := 6
	//newsCache.DBGetter = func() string {
	//	log.Println("get form db")
	//	newsModel := lib.NewNewsModel()
	//	lib.Gorm.Table("mynews").Where("id=?",newsID).Find(newsModel)
	//	b, _ := json.Marshal(newsModel)
	//	return string(b)
	//}
	//
	//newsCache.DBGetter = func() string {
	//	log.Println("get from db")
	//	return "data from db"
	//}
	//fmt.Println(newsCache.GetCache("news123").(*lib.NewsModel).NewsTitle)   // get from news

	r := gin.New()
	r.Handle("GET", "/news/:id", func(context *gin.Context) {
		//1. 从对象池 获取新闻缓存 对象
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)

		//2. 获取参数, 设置DBGetter
		newsID := context.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID) //一旦缓存没有，则需要从哪里去取

		//3.取缓存输出(一旦没有，上面的DBGetter会被调用)
		context.Header("Content-type", "application/json")
		context.String(200, newsCache.GetCache("news"+newsID).(string))
	})

	r.Run(":8080")
}
