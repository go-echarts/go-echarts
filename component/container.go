package component

import (
	"github.com/go-echarts/go-echarts/v2/event"
	"github.com/go-echarts/go-echarts/v2/util"
)

// Container Each chart map to a Container and vice versa , it is 1:1, they hold each other
type Container struct {

	// Width of component
	Width string

	// Height of component
	Height string

	// BackgroundColor of component
	BackgroundColor string

	// Chart unique ID, as same as the Container Id
	ChartID string

	// Theme of chart
	Theme string

	// 1:1
	Chart Chart

	// the abstract of all the event which binds to an echarts instance
	Events []*event.Event `json:"-"`

	// Inline styles
	// Style
}

func (c *Container) AddEvent(event *event.Event) {
	c.Events = append(c.Events, event)
}

func NewContainer(chart Chart) *Container {
	return &Container{
		Width:   "900px",
		Height:  "500px",
		ChartID: util.GenerateUniqueID(),
		Theme:   "white",
		Chart:   chart,
	}

}
