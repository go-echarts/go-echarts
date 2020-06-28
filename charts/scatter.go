package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Scatter represents a scatter chart.
type Scatter struct {
	RectChart
}

func (Scatter) Type() string { return types.ChartScatter }

// NewScatter creates a new scatter chart.
func NewScatter() *Scatter {
	chart := &Scatter{}
	chart.initBaseConfiguration()
	chart.Renderer = render.NewChartRender(chart, chart.Validate)
	chart.HasXYAxis = true
	return chart
}

// SetXAxis adds the X axis.
func (c *Scatter) SetXAxis(x interface{}) *Scatter {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *Scatter) AddSeries(name string, data []opts.ScatterData, opts ...SeriesOpts) *Scatter {
	series := SingleSeries{Name: name, Type: types.ChartScatter, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
