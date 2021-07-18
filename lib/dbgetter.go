package lib

import (
	"Go_Redis/gedis"
	"encoding/json"
	"log"
)

func NewsDBGetter(id string) gedis.DBGettFuc {
	return func() string {
		log.Println("get from db")
		newsModel := NewNewsModel()
		Gorm.Table("mynews").Where("id=?", id).Find(newsModel)
		b, _ := json.Marshal(newsModel)
		return string(b)
	}
}
