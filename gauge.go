package goecharts

import (
	"io"
)

type Gauge struct {
	BaseOpts
	Series
}

// 工厂函数，生成 `Gauge` 实例
func NewGauge(routers ...HTTPRouter) *Gauge {
	chart := new(Gauge)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Gauge) Add(name string, data map[string]interface{}, options ...interface{}) *Gauge {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "gauge", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Gauge) SetGlobalConfig(options ...interface{}) *Gauge {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Gauge) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", w...)
}
