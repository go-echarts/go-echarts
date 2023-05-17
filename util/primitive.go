package util

import "github.com/go-echarts/go-echarts/v2/primitive"

func Strconv(data ...string) []primitive.String {
	var ls []primitive.String
	for _, e := range data {

		ls = append(ls, primitive.StringOf(e))
	}
	return ls
}
