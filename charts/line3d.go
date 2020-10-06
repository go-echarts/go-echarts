package charts

import (
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Line3D represents a 3D line chart.
type Line3D struct {
	Chart3D
}

// Type returns the chart type.
func (Line3D) Type() string { return types.ChartLine3D }

// NewLine3D creates a new 3D line chart.
func NewLine3D() *Line3D {
	c := &Line3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Line3D) AddXYAxis(xAxis, yAxis interface{}) *Line3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Line3D) AddZAxis(name string, zAxis interface{}, options ...SeriesOpts) *Line3D {
	c.addZAxis(types.ChartLine3D, name, zAxis, options...)
	return c
}
