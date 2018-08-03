package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ledboot/Lynx/router/v1"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/shortUrl", v1.GetUrl)
	}

	r.GET("/ws",v1.WsHandler)

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "chat.html", gin.H{})
	})

	gin.SetMode(gin.DebugMode)

	return r
}
