package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
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

// AddSeries adds the new series.
func (c *Line3D) AddSeries(name string, data []opts.Chart3DData, options ...SeriesOpts) *Line3D {
	c.addSeries(types.ChartLine3D, name, data, options...)
	return c
}
