package charts

import "github.com/go-echarts/go-echarts/types"

// EffectScatter represents an effect scatter chart.
type EffectScatter struct {
	RectChart
}

func (EffectScatter) Type() string { return types.ChartEffectScatter }

// EffectScatterChartOpts is the option set for an effect scatter chart.
type EffectScatterChartOpts struct {
	XAxisIndex int
	YAxisIndex int
}

// NewEffectScatter creates a new effect scatter chart.
func NewEffectScatter() *EffectScatter {
	chart := new(EffectScatter)
	chart.initBaseConfiguration()
	chart.initXYAxis()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *EffectScatter) AddYAxis(name string, yAxis interface{}, opts ...SeriesOpts) *EffectScatter {
	series := SingleSeries{Name: name, Type: types.ChartEffectScatter, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
