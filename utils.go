package goecharts

import (
	"log"
)

// 异常检测
func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 翻转 string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
