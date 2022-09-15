package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Funnel represents a funnel chart.
type Funnel struct {
	BaseConfiguration
	BaseActions
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
func (c *Funnel) AddSeries(name string, data []opts.FunnelData, options ...SeriesOpts) *Funnel {
	series := SingleSeries{Name: name, Type: types.ChartFunnel, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Funnel instance.
func (c *Funnel) SetGlobalOptions(options ...GlobalOpts) *Funnel {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions sets actions for the Gauge instance.
func (c *Funnel) SetDispatchActions(actions ...GlobalActions) *Funnel {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate validates the given configuration.
func (c *Funnel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
