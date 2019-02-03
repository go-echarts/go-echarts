package charts

import (
	"io"
)

type Funnel struct {
	BaseOpts
	Series
}

func (Funnel) chartType() string { return "funnel" }

func NewFunnel(routers ...HTTPRouter) *Funnel {
	chart := new(Funnel)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Funnel) Add(name string, data map[string]interface{}, options ...seriesOptser) *Funnel {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "funnel", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Funnel) SetGlobalOptions(options ...globalOptser) *Funnel {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Funnel) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Funnel) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
