package lib

import (
	"regexp"
	"strconv"
)

var tenToAny = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func initArray() {
	condition := '0'
	var conditonStr string
	for condition <= 'z' {
		conditonStr = string(condition)

		matched, _ := regexp.MatchString(`[\d]|[A-Za-z]`, conditonStr)
		if matched {
			tenToAny = append(tenToAny, conditonStr)
		}
		condition++
	}
}

func GetShortCode(num, n int64) string {

	remainder_str := ""
	if num < 0 {
		return remainder_str
	}
	var remainder int64
	for num != 0 {
		remainder = num % n
		if remainder < 62 && remainder > 10 {
			remainder_str = tenToAny[remainder] + remainder_str
		} else {
			remainder_str = strconv.FormatInt(int64(remainder), 10) + remainder_str
		}
		num = num / n
	}

	return remainder_str
}
