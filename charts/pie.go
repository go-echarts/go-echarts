package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/datatypes"
)

// Pie represents a pie chart.
type Pie struct {
	BaseOpts
	MultiSeries
}

func (Pie) Type() string { return ChartType.Pie }

// NewPie creates a new gauge chart.
func NewPie(routers ...RouterOpts) *Pie {
	chart := new(Pie)
	chart.initBaseOpts(routers...)
	return chart
}

// Add adds new data sets.
func (c *Pie) Add(name string, data map[string]interface{}, fns ...SeriesOpts) *Pie {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: ChartType.Pie, Data: nvs}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.setColor(options...)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Pie) SetGlobalOptions(options ...GlobalOptser) *Pie {
	c.BaseOpts.setBaseGlobalOptions(options...)
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
