package goecharts

const (
	echartsJS = "echarts.min.js"
	bulmaCSS  = "bulma.min.css"
)

// Name-Value 数据项
type nameValueItem struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
