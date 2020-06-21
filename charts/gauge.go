package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/types"
)

// Gauge represents a gauge chart.
type Gauge struct {
	BaseConfiguration
	MultiSeries
}

func (Gauge) Type() string { return types.ChartGauge }

// NewGauge creates a new gauge chart.
func NewGauge() *Gauge {
	chart := &Gauge{}
	chart.initBaseConfiguration()
	return chart
}

// Add adds new data sets.
func (c *Gauge) AddSeries(name string, data map[string]interface{}, opts ...SeriesOpts) *Gauge {
	nvs := make([]types.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, types.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: types.ChartGauge, Data: nvs}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Gauge instance.
func (c *Gauge) SetGlobalOptions(opts ...GlobalOpts) *Gauge {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Gauge) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.Validate()
	return renderToWriter(c, "chart", []string{}, w...)
}
