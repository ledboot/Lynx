package main

import (
	"github.com/ledboot/Lynx/router"
	"github.com/ledboot/Lynx/models"
)

func main() {
	models.SetupEngine()
	models.Sync()
	r := router.SetupRouter()
	r.Run(":8091")
}
