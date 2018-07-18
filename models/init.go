package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/gomodule/redigo/redis"
	"time"
	"fmt"
)

var engine *xorm.Engine
var pool *redis.Pool

func SetupEngine() {
	setupMysql()
	setupRedis()
}

func GetRedis() redis.Conn {
	if pool == nil {
		fmt.Println("redis can not use!")
		return nil
	}
	return pool.Get()
}

func EnableRedis() bool {
	return pool != nil
}

func setupMysql() {
	fmt.Println("setup mysql...")
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:32768)/lynx?charset=utf8")
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(0)
	engine.SetMaxOpenConns(5)
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetTableMapper(core.GonicMapper{})
	engine.SetColumnMapper(core.GonicMapper{})
}

func setupRedis() {
	fmt.Println("setup redis...")
	pool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   6,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:32772")
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", "root"); err != nil {
				conn.Close()
				return nil, err
			}
			if _, err := conn.Do("select", 0); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
	}
}

func Sync() {
	engine.Sync2(new(Indexs))
	engine.Query("alter table indexs auto_increment =1000")
}
