package charts

import "github.com/go-echarts/go-echarts/types"

// Kline represents a kline chart.
type Kline struct {
	RectChart
}

func (Kline) Type() string { return types.ChartKline }

// NewKLine creates a new kline chart.
func NewKLine() *Kline {
	chart := &Kline{}
	chart.initBaseConfiguration()
	chart.initXYAxis()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Kline) SetXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Kline) AddSeries(name string, yAxis interface{}, opts ...SeriesOpts) *Kline {
	series := SingleSeries{Name: name, Type: types.ChartKline, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
