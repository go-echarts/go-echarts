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
