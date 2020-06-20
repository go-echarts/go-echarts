package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/types"
)

// Liquid represents a liquid chart.
type Liquid struct {
	BaseConfiguration
	MultiSeries
}

func (Liquid) Type() string { return types.ChartLiquid }

// NewLiquid creates a new liquid chart.
func NewLiquid() *Liquid {
	chart := &Liquid{}
	chart.JSAssets.Add("echarts-liquidfill.min.js")
	return chart
}

// Add adds new data sets.
func (c *Liquid) Add(name string, data interface{}, opts ...SeriesOpts) *Liquid {
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

func (c *Liquid) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Liquid) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.Validate()
	return renderToWriter(c, "chart", []string{`"outline":{"show":false},?`, `"waveAnimation":false,?`}, w...)
}
