package v1

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"github.com/ledboot/Lynx/lib"
	"github.com/ledboot/Lynx/database"
	"github.com/ledboot/Lynx/models"
	"net/http"
)

func GetUrl(context *gin.Context) {
	url := context.Query("url")
	fmt.Println("query url ->", url)
	result, _ := database.GetEngine().Query("select max(id) from indexs")
	maxIdStr := string(result[0]["max(id)"])
	var maxId int64 = 1000
	if maxIdStr != "" {
		maxId, _ = strconv.ParseInt(maxIdStr, 10, 64)
		maxId++
	}
	fmt.Println("maxid ->", maxId)
	keyword := lib.GetShortCode(maxId, 62)

	database.GetEngine().Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})

	response := "https://t.cn/" + keyword
	fmt.Println(response)
	context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})
}
