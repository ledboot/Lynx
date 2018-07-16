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

	if url == "" {
		context.JSON(http.StatusOK, gin.H{"code": "400", "message": "request url is nil", "data": ""})
		return
	}
	fmt.Println("query url ->", url)
	response := "https://t.cn/"

	if key, find := findKey(url); find {
		response += key
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

		response += keyword
		mysql.GetEngine().Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})
		if wapperRedis.Enable() {
			wapperRedis.GetEngine().Do("set", url, keyword, "ex 10")
		}
		context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})
	}

}

func findKey(url string) (string, bool) {
	var flag bool
	var key string
	var indexs models.Indexs
	if wapperRedis.Enable() {
		reply, err := wapperRedis.GetEngine().Do("get", url)
		if err == nil {
			code, _ := redis.String(reply, err)
			if code != "" {
				fmt.Println("get key from redis ->", code)
				key = code
				flag = true
			} else {
				mysql.GetEngine().Where("url = ?", url).Get(&indexs)
				if &indexs != nil && indexs.Id != 0 && indexs.Keyword != "" {
					key = indexs.Keyword
					flag = true
					if wapperRedis.Enable() {
						reply, error := wapperRedis.GetEngine().Do("set", url, key, "EX", "43200")
						fmt.Println(reply, error)
					}
				}
			}
		}
	} else {
		mysql.GetEngine().Where("url = ?", url).Get(&indexs)
		if &indexs != nil && indexs.Keyword != "" {
			key = indexs.Keyword
			flag = true
			if wapperRedis.Enable() {
				reply, error := wapperRedis.GetEngine().Do("set", url, key, "EX", "43200")
				fmt.Println(reply, error)
			}
		}

	}
	return key, flag
}
