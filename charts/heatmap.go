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

// SetXAxis adds the X axis.
func (c *HeatMap) SetXAxis(x interface{}) *HeatMap {
	c.xAxisData = x
	return c
}

// AddSeries adds the Y axis.
func (c *HeatMap) AddSeries(name string, yAxis interface{}, opts ...SeriesOpts) *HeatMap {
	series := SingleSeries{Name: name, Type: types.ChartHeatMap, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
