package util

import "github.com/go-echarts/go-echarts/v2/primitive"

func StrConv(str string) primitive.String {
	return primitive.StringOf(str)
}

func StringsConv(data ...string) []primitive.String {
	var ls []primitive.String
	for _, e := range data {

		ls = append(ls, primitive.StringOf(e))
	}
	return ls
}

func FloatConv(f float32) primitive.Float {
	return primitive.FloatOf(f)
}

func FloatsConv(data ...float32) []primitive.Float {
	var ls []primitive.Float
	for _, e := range data {

		ls = append(ls, primitive.FloatOf(e))
	}
	return ls
}

func IntConv(i int) primitive.Int {
	return primitive.IntOf(i)
}

func IntsConv(data ...int) []primitive.Int {
	var ls []primitive.Int
	for _, e := range data {

		ls = append(ls, primitive.IntOf(e))
	}
	return ls
}
