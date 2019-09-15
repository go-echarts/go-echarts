package charts

import (
	"github.com/go-echarts/go-echarts/datatypes"
	"io"
)

// Funnel represents a funnel chart.
type Funnel struct {
	BaseOpts
	Series
}

func (Funnel) chartType() string { return ChartType.Funnel }

// NewFunnel creates a new funnel chart.
func NewFunnel(routers ...RouterOpts) *Funnel {
	chart := new(Funnel)
	chart.initBaseOpts(routers...)
	return chart
}

// Add adds new data sets.
func (c *Funnel) Add(name string, data map[string]interface{}, options ...seriesOptser) *Funnel {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: ChartType.Funnel, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

// SetGlobalOptions sets options for the Funnel instance.
func (c *Funnel) SetGlobalOptions(options ...globalOptser) *Funnel {
	c.BaseOpts.setBaseGlobalOptions(options...)
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
