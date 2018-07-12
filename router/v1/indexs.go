package v1

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/ledboot/Lynx/models"
	"github.com/ledboot/Lynx/database/mysql"
	wapperRedis "github.com/ledboot/Lynx/database/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/ledboot/Lynx/lib"
	"net/http"
	"strconv"
)

func GetUrl(context *gin.Context) {
	url := context.Query("url")
	fmt.Println("query url ->", url)
	response := "https://t.cn/"
	findKey := false
	if wapperRedis.Enable() {
		reply, err := wapperRedis.GetEngine().Do("dump", url)
		if err == nil {
			fmt.Println(reply, err)
			code, _ := redis.String(reply, err)
			if code != "" {
				response += code
				findKey = true
			}
		}
	} else {
		var indexs models.Indexs
		mysql.GetEngine().Where("url = ?", url).Get(&indexs)
		if &indexs != nil {
			response += indexs.Keyword
		}
		findKey = true
	}
	if findKey {
		context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})
	} else {
		result, _ := mysql.GetEngine().Query("select max(id) from indexs")
		maxIdStr := string(result[0]["max(id)"])
		var maxId int64 = 1000
		if maxIdStr != "" {
			maxId, _ = strconv.ParseInt(maxIdStr, 10, 64)
			maxId++
		}
		fmt.Println("maxid ->", maxId)
		keyword := lib.GetShortCode(maxId, 62)

		mysql.GetEngine().Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})
		if wapperRedis.Enable() {
			wapperRedis.GetEngine().Do("set", url, keyword)
		}
		fmt.Println(response)
		context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})
	}

}
