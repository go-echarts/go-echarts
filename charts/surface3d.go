package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Surface3D represents a 3D surface chart.
type Surface3D struct {
	Chart3D
}

// Type returns the chart type.
func (Surface3D) Type() string { return types.ChartSurface3D }

// NewSurface3D creates a new 3d surface chart.
func NewSurface3D() *Surface3D {
	c := &Surface3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

// AddSeries adds the new series.
func (c *Surface3D) AddSeries(name string, data []opts.Chart3DData, options ...SeriesOpts) *Surface3D {
	c.addSeries(types.ChartScatter3D, name, data, options...)
	return c
}
