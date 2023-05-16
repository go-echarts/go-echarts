package components

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/util"
)

type Container struct {

	// Width of canvas
	Width string `default:"900px"`

	// Height of canvas
	Height string `default:"500px"`

	// BackgroundColor of canvas
	BackgroundColor string

	// Chart unique ID, as same as the Container Id
	ChartID string

	// Theme of chart
	Theme string

	// 1:1
	Chart interface{}
}

func NewDefaultContainer(chart interface{}) *Container {
	util.ConfigPrettier(chart)
	return &Container{
		Width:   "900px",
		Height:  "500px",
		ChartID: opts.GenUUID(),
		Theme:   "white",
		Chart:   chart,
	}

}
