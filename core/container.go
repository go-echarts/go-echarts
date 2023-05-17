package core

import (
	"github.com/go-echarts/go-echarts/v2/util"
)

// Container Each chart belongs to a Container, it is 1:1
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

func NewContainer(chart interface{}) *Container {
	return &Container{
		Width:   "900px",
		Height:  "500px",
		ChartID: util.GenerateUniqueID(),
		Theme:   "white",
		Chart:   chart,
	}

}
