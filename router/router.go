package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ledboot/Lynx/router/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/shortUrl", v1.GetUrl)
	}

	gin.SetMode(gin.DebugMode)

	return r
}
