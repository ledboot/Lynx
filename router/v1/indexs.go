package v1

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/ledboot/Lynx/models"
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
		maxIdStr := models.FindIndexMaxId()
		var maxId int64 = 1000
		if maxIdStr != "" {
			maxId, _ = strconv.ParseInt(maxIdStr, 10, 64)
			maxId++
		}
		fmt.Println("maxid ->", maxId)
		keyword := lib.GetShortCode(maxId, 62)

		response += keyword
		models.Insert(models.Indexs{Url: url, Keyword: keyword, UrlType: "system"})
		redisSaveKey(url, keyword)
		context.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": response})
	}

}

func findKey(url string) (string, bool) {
	var flag bool
	var key string
	if models.EnableRedis() {
		code, err := redisGetValue(key)
		if err == nil && code != "" {
			fmt.Println("get key from redis ->", code)
			key = code
			flag = true
		}
	}
	if !false {
		indexs, find := models.GetIndexByUrl(url)
		if find {
			key = indexs.Keyword
			flag = true
			redisSaveKey(url, key)
		}
	}
	return key, flag
}

func redisSaveKey(url, key string) {
	if models.EnableRedis() {
		reply, error := models.GetRedis().Do("set", url, key, "EX", "43200")
		fmt.Println(reply, error)
	}
}

func redisGetValue(key string) (string, error) {
	return redis.String(models.GetRedis().Do("get", key))
}
