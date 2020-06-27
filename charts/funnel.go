package charts

import (
	"io"

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
	chart := &Funnel{}
	chart.initBaseConfiguration()
	return chart
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

// Render renders the chart and writes the output to given writer.
func (c *Funnel) Render(w io.Writer) error {
	c.Validate()
	return render.ChartRender(c, w)
}
