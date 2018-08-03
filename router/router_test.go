package router

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/ledboot/Lynx/models"
)

func TestShortUrl(t *testing.T) {
	models.SetupEngine()
	models.Sync()
	router := SetupRouter()
	w := httptest.NewRecorder()
	params := "https://www.baidu.com"
	req, _ := http.NewRequest("GET", `/api/v1/shortUrl?url=`+params, nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Code, w.Body)
}
