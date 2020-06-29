package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Funnel represents a funnel chart.
type Funnel struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (Funnel) Type() string { return types.ChartFunnel }

// NewFunnel creates a new funnel chart.
func NewFunnel() *Funnel {
	c := &Funnel{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *Funnel) AddSeries(name string, data []opts.FunnelData, opts ...SeriesOpts) *Funnel {
	series := SingleSeries{Name: name, Type: types.ChartFunnel, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Funnel instance.
func (c *Funnel) SetGlobalOptions(opts ...GlobalOpts) *Funnel {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

// Validate
func (c *Funnel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
