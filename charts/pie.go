package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/types"
)

// Pie represents a pie chart.
type Pie struct {
	BaseConfiguration
	MultiSeries
}

func (Pie) Type() string { return types.ChartPie }

// NewPie creates a new gauge chart.
func NewPie() *Pie {
	chart := new(Pie)
	chart.initBaseConfiguration()
	return chart
}

// Add adds new data sets.
func (c *Pie) Add(name string, data map[string]interface{}, opts ...SeriesOpts) *Pie {
	nvs := make([]types.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, types.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: types.ChartPie, Data: nvs}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Pie) SetGlobalOptions(opts ...GlobalOpts) *Pie {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Pie) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Pie) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
