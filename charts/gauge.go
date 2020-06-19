package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/datatypes"
)

// Gauge represents a gauge chart.
type Gauge struct {
	BaseConfiguration
	MultiSeries
}

func (Gauge) Type() string { return ChartType.Gauge }

// NewGauge creates a new gauge chart.
func NewGauge(routers ...RouterOpts) *Gauge {
	chart := new(Gauge)
	chart.initBaseConfiguration(routers...)
	return chart
}

// Add adds new data sets.
func (c *Gauge) Add(name string, data map[string]interface{}, opts ...SeriesOpts) *Gauge {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: ChartType.Gauge, Data: nvs}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Gauge instance.
func (c *Gauge) SetGlobalOptions(opts ...GlobalOpts) *Gauge {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Gauge) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
