package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Liquid represents a liquid chart.
type Liquid struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (Liquid) Type() string { return types.ChartLiquid }

// NewLiquid creates a new liquid chart.
func NewLiquid() *Liquid {
	c := &Liquid{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.JSAssets.Add("echarts-liquidfill.min.js")
	return c
}

// AddSeries adds new data sets.
func (c *Liquid) AddSeries(name string, data []opts.LiquidData, opts ...SeriesOpts) *Liquid {
	series := SingleSeries{Name: name, Type: types.ChartLiquid, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Liquid instance.
func (c *Liquid) SetGlobalOptions(opts ...GlobalOpts) *Liquid {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

// Validate
func (c *Liquid) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
