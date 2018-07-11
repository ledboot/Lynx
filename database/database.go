package database

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ledboot/Lynx/models"
)

var _engine *xorm.Engine

func ConfigDb() {
	engine, error := xorm.NewEngine("mysql", "root:root@tcp(localhost:32768)/lynx?charset=utf8")
	if error != nil {
		panic(error)
	}
	engine.SetMaxIdleConns(0)
	engine.SetMaxOpenConns(5)
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetTableMapper(core.GonicMapper{})
	engine.SetColumnMapper(core.GonicMapper{})
	_engine = engine
}

func GetEngine() *xorm.Engine {
	if _engine != nil {
		return _engine
	} else {
		panic("please config db first")
	}
}

func Sync() {
	_engine.Sync2(new(models.Indexs))
	_engine.Query("alter table indexs auto_increment =1000")
}
