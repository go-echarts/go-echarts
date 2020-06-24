package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

// Line represents a line chart.
type Line struct {
	RectChart
}

// Type returns the chart type.
func (Line) Type() string { return types.ChartLine }

// NewLine creates a new line chart.
func NewLine() *Line {
	chart := &Line{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// SetXAxis adds the X axis.
func (c *Line) SetXAxis(x interface{}) *Line {
	c.xAxisData = x
	return c
}

// AddSeries adds the Y axis.
func (c *Line) AddSeries(name string, data []opts.LineChartItem, opts ...SeriesOpts) *Line {
	series := SingleSeries{Name: name, Type: types.ChartLine, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
