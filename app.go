package main

import (
	"github.com/ledboot/Lynx/router"
	"github.com/ledboot/Lynx/database/mysql"
	"github.com/ledboot/Lynx/database/redis"
)

func main() {
	mysql.Config()
	mysql.Sync()
	redis.Config()

	r := router.SetupRouter()
	r.Run(":8090")
}
