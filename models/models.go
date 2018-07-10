package models

import "time"

type Indexs struct {
	Id        int64
	Url       string    `xorm:"varchar(255) notnull"`
	Keyword   string    `xorm:"varchar(20) notnull"`
	UrlType   string    `xorm:"varchar(10) notnull"`
	CreateaAt time.Time `xorm:"DateTime notnull created"`
	UpdateAt  time.Time `xorm:"DateTime notnull updated"`
}
