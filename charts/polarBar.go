package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// PolarBar represents a polar bar chart.
type PolarBar struct {
	RectChart
}

// Type returns the chart type.
func (PolarBar) Type() string { return types.ChartBar }

// NewPolarBar creates a new bar chart instance.
func NewPolarBar() *PolarBar {
	c := &PolarBar{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds the new series.
func (c *PolarBar) AddSeries(name string, data []opts.BarData, options ...SeriesOpts) *PolarBar {
	series := SingleSeries{Name: name, Type: types.ChartBar, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *PolarBar) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
