package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Scatter3D represents a 3D scatter chart.
type Scatter3D struct {
	Chart3D
}

// Type returns the chart type.
func (Scatter3D) Type() string { return types.ChartScatter3D }

// NewScatter3D creates a new 3D scatter chart.
func NewScatter3D() *Scatter3D {
	c := &Scatter3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

// AddSeries adds the new series.
func (c *Scatter3D) AddSeries(name string, data []opts.Chart3DData, options ...SeriesOpts) *Scatter3D {
	c.addSeries(types.ChartScatter3D, name, data, options...)
	return c
}
