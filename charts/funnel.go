package charts

import (
	"io"
)

type Funnel struct {
	BaseOpts
	Series
}

// 工厂函数，生成 `Funnel` 实例
func NewFunnel(routers ...HTTPRouter) *Funnel {
	chart := new(Funnel)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Funnel) Add(name string, data map[string]interface{}, options ...interface{}) *Funnel {
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

func (c *Funnel) SetGlobalConfig(options ...interface{}) *Funnel {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Funnel) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Funnel) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", w...)
}
