package core

import (
	"github.com/go-echarts/go-echarts/v2/util"
)

// Container Each chart belongs to a Container, it is 1:1
type Container struct {

	// Width of container
	Width string

	// Height of container
	Height string

	// BackgroundColor of container
	BackgroundColor string

	// Chart unique ID, as same as the Container Id
	ChartID string

	// Theme of chart
	Theme string

	// 1:1
	Chart interface{}

	Events []*Event `json:"-"`

	// Inline styles
	// Style
}

func (c *Container) AddEvent(event *Event) {
	c.Events = append(c.Events, event)
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
