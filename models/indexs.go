package models

import (
	"time"
)

type Indexs struct {
	Id        int64
	Url       string    `xorm:"varchar(255) notnull unique"`
	Keyword   string    `xorm:"varchar(20) notnull unique"`
	UrlType   string    `xorm:"varchar(10) notnull"`
	CreateaAt time.Time `xorm:"DateTime notnull created"`
	UpdateAt  time.Time `xorm:"DateTime notnull updated"`
}

func GetIndexByUrl(url string) (Indexs, bool) {
	var indexs Indexs
	var find bool
	engine.Where("url = ?", url).Get(&indexs)
	if indexs.Id != 0 {
		find = true
	}
	return indexs, find
}

func FindIndexMaxId() string {
	var maxIdStr string
	result, _ := engine.Query("select max(id) from indexs")
	if len(result) > 0 {
		maxIdStr = string(result[0]["max(id)"])
	}
	return maxIdStr
}

func Insert(indexs Indexs) {
	engine.Insert(indexs)
}
