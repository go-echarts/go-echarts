package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

// EffectScatter represents an effect scatter chart.
type EffectScatter struct {
	RectChart
}

func (EffectScatter) Type() string { return types.ChartEffectScatter }

// NewEffectScatter creates a new effect scatter chart.
func NewEffectScatter() *EffectScatter {
	chart := &EffectScatter{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *EffectScatter) SetXAxis(x interface{}) *EffectScatter {
	c.xAxisData = x
	return c
}

// AddYAxis adds the Y axis.
func (c *EffectScatter) AddSeries(name string, data []opts.EffectScatterData, opts ...SeriesOpts) *EffectScatter {
	series := SingleSeries{Name: name, Type: types.ChartEffectScatter, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
