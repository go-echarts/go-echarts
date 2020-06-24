package charts

import "github.com/go-echarts/go-echarts/types"

// BoxPlot represents a boxplot chart.
type BoxPlot struct {
	RectChart
}

func (BoxPlot) Type() string { return types.ChartBoxPlot }

// NewBoxPlot creates a new boxplot chart.
func NewBoxPlot() *BoxPlot {
	chart := &BoxPlot{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// SetXAxis adds the X axis.
func (c *BoxPlot) SetXAxis(x interface{}) *BoxPlot {
	c.xAxisData = x
	return c
}

// AddSeries adds the Y axis.
func (c *BoxPlot) AddSeries(name string, yAxis interface{}, opts ...SeriesOpts) *BoxPlot {
	series := SingleSeries{Name: name, Type: types.ChartBoxPlot, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
