package goecharts

const (
	echartsJS = "echarts.min.js"
	bulmaCSS  = "bulma.min.css"

	geoFormatter = `function (params) {return params.name + ' : ' + params.value[2];}`
)

// Name-Value 数据项
type nameValueItem struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
