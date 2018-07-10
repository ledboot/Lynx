package main

import (
	"github.com/ledboot/Lynx/database"
	"github.com/ledboot/Lynx/router"
)

func main() {
	database.ConfigDb()
	database.Sync()

	r := router.SetupRouter()
	r.Run(":8090")
}
