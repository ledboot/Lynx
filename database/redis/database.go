package redis

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"time"
)

var _poll *redis.Pool

func Config() {
	_poll = &redis.Pool{
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
	//_poll.Stats()
	//
	//client, error := redis.Dial("tcp", "127.0.0.1:32772")
	//defer client.Close()
	//if error != nil {
	//	fmt.Println("redis connect fail !", error.Error())
	//	return
	//}
	//_client = client
}

func Enable() bool {
	return _poll != nil
}

func GetEngine() redis.Conn {
	if _poll == nil {
		fmt.Println("redis can not use!")
		return nil
	}
	return _poll.Get()
}
