package lib

import (
	"testing"
	"fmt"
)

func TestGetShortCode(t *testing.T) {
	//var i int64 = 1
	//for ; i < 5000000; i++ {
	//	code := GetShortCode(i, 62)
	//	fmt.Println(code)
	//}

	code := GetShortCode(1000, 62)
	fmt.Println(code)
}
