package charts

import (
	"github.com/go-echarts/go-echarts/types"
	"io"
)

// Funnel represents a funnel chart.
type Funnel struct {
	BaseConfiguration
	MultiSeries
}

func (Funnel) Type() string { return types.ChartFunnel }

// NewFunnel creates a new funnel chart.
func NewFunnel() *Funnel {
	chart := new(Funnel)
	//chart.initBaseOpts(routers...)
	return chart
}

// Add adds new data sets.
func (c *Funnel) Add(name string, data map[string]interface{}, opts ...SeriesOpts) *Funnel {
	nvs := make([]types.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, types.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: types.ChartFunnel, Data: nvs}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Funnel instance.
func (c *Funnel) SetGlobalOptions(opts ...GlobalOpts) *Funnel {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Funnel) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Funnel) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
