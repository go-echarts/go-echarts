package charts

import (
	"github.com/go-echarts/go-echarts/opts"
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
func (c *Gauge) AddSeries(name string, data []opts.GaugeChartItem, opts ...SeriesOpts) *Gauge {
	series := SingleSeries{Name: name, Type: types.ChartGauge, Data: data}
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
func (c *Gauge) Render(w io.Writer) error {
	c.Validate()
	return renderToWriter(c, ModChart, w)
}
