package charts

import "github.com/go-echarts/go-echarts/types"

// BoxPlot represents a boxplot chart.
type BoxPlot struct {
	RectChart
}

func (BoxPlot) Type() string { return types.ChartBoxPlot }

// NewBoxPlot creates a new boxplot chart.
func NewBoxPlot() *BoxPlot {
	chart := new(BoxPlot)
	chart.initBaseConfiguration()
	chart.initXYAxis()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *BoxPlot) AddXAxis(xAxis interface{}) *BoxPlot {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *BoxPlot) AddYAxis(name string, yAxis interface{}, opts ...SeriesOpts) *BoxPlot {
	series := SingleSeries{Name: name, Type: types.ChartBoxPlot, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
