package charts

import (
	"io"
)

// Liquid represents a liquid chart.
type Liquid struct {
	BaseConfiguration
	MultiSeries
}

func (Liquid) Type() string { return ChartType.Liquid }

// NewLiquid creates a new liquid chart.
func NewLiquid() *Liquid {
	chart := new(Liquid)
	chart.JSAssets.Add("echarts-liquidfill.min.js")
	return chart
}

// Add adds new data sets.
func (c *Liquid) Add(name string, data interface{}, opts ...SeriesOpts) *Liquid {
	series := SingleSeries{Name: name, Type: ChartType.Liquid, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Liquid instance.
func (c *Liquid) SetGlobalOptions(opts ...GlobalOpts) *Liquid {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Liquid) validateOpts() {
	c.ValidateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Liquid) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{`"outline":{"show":false},?`, `"waveAnimation":false,?`}, w...)
}
