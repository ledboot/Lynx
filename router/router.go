package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ledboot/Lynx/database"
	"strconv"
	"fmt"
	"github.com/ledboot/Lynx/lib"
	"github.com/ledboot/Lynx/models"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	//gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(200, "hello~")
	})

	v1 := r.Group("v1")
	{
		v1.GET("/shortUrl", func(context *gin.Context) {
			url := context.Param("url")
			fmt.Println(url)
			result, _ := database.GetEngine().Query("select max(id) from indexs")
			maxIdStr := string(result[0]["max(id)"])
			var maxId int64 = 1
			if maxIdStr != "" {
				maxId, _ = strconv.ParseInt(maxIdStr, 10, 64)
				maxId++
			}
			keyword := lib.GetShortCode(maxId, 62)

			database.GetEngine().Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})

			response := "https://t.cn/" + keyword
			fmt.Println(response)
			context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})

		})
	}

	gin.SetMode(gin.DebugMode)
	return r
}

func shortUrl(c *gin.Context) {
	url := c.Params.ByName("url")
	fmt.Println(url)
	result, _ := database.GetEngine().Query("select max(id) from indexs")
	maxIdStr := string(result[0]["max(id)"])
	var maxId int64 = 1
	if maxIdStr != "" {
		maxId, _ = strconv.ParseInt(maxIdStr, 10, 64)
		maxId++
	}
	keyword := lib.GetShortCode(maxId, 64)

	database.GetEngine().Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})

	response := "https://t.cn/" + keyword
	c.String(200, response)
}
