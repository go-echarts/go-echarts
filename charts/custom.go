package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Custom represents a custom chart.
type Custom struct {
	RectChart
}

// Type returns the chart type.
func (*Custom) Type() string { return types.ChartCustom }

// NewCustom creates a new Custom chart.
func NewCustom() *Custom {
	c := &Custom{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *Custom) SetXAxis(x interface{}) *Custom {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *Custom) AddSeries(name string, data []opts.CustomData, options ...SeriesOpts) *Custom {
	series := SingleSeries{Name: name, Type: types.ChartCustom, Data: data}
	series.InitSeriesDefaultOpts(c.BaseConfiguration)
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *Custom) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
