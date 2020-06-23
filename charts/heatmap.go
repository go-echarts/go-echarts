package charts

import "github.com/go-echarts/go-echarts/types"

// HeatMap represents a heatmap chart.
type HeatMap struct {
	RectChart
}

func (HeatMap) Type() string { return types.ChartHeatMap }

// NewHeatMap creates a new heatmap chart.
func NewHeatMap() *HeatMap {
	chart := &HeatMap{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *HeatMap) AddXAxis(xAxis interface{}) *HeatMap {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *HeatMap) AddYAxis(name string, yAxis interface{}, opts ...SeriesOpts) *HeatMap {
	series := SingleSeries{Name: name, Type: types.ChartHeatMap, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
