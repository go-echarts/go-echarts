package charts

import (
	"io"
)

type Gauge struct {
	BaseOpts
	Series
}

func (Gauge) chartType() string { return "gauge" }

func NewGauge(routers ...HTTPRouter) *Gauge {
	chart := new(Gauge)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Gauge) Add(name string, data map[string]interface{}, options ...seriesOptser) *Gauge {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "gauge", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

func (c *Gauge) SetGlobalOptions(options ...globalOptser) *Gauge {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Gauge) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
