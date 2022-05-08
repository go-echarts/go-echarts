package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Line represents a line chart.
type Line struct {
	RectChart
}

// Type returns the chart type.
func (Line) Type() string { return types.ChartLine }

// NewLine creates a new line chart.
func NewLine() *Line {
	c := &Line{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *Line) SetXAxis(x interface{}) *Line {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *Line) AddSeries(name string, data []opts.LineData, options ...SeriesOpts) *Line {
	series := SingleSeries{Name: name, Type: types.ChartLine, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *Line) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
