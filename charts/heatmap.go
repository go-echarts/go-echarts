package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

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

// AddSeries adds the new series.
func (c *HeatMap) AddSeries(name string, data []opts.HeatMapData, opts ...SeriesOpts) *HeatMap {
	series := SingleSeries{Name: name, Type: types.ChartHeatMap, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
