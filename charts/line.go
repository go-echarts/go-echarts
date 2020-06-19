package charts

import "github.com/go-echarts/go-echarts/types"

// Line represents a line chart.
type Line struct {
	RectChart
}

func (Line) Type() string { return types.ChartLine }

// NewLine creates a new line chart.
func NewLine() *Line {
	chart := new(Line)
	chart.initBaseConfiguration()
	chart.initXYAxis()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Line) AddXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Line) AddYAxis(name string, yAxis interface{}, opts ...SeriesOpts) *Line {
	series := SingleSeries{Name: name, Type: types.ChartLine, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
