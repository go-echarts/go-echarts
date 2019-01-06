package goecharts

import (
	"io"
)

type Pie struct {
	BaseOpts
	Series
}

// 工厂函数，生成 `Pie` 实例
func NewPie(routers ...HTTPRouter) *Pie {
	chart := new(Pie)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Pie) Add(name string, data map[string]interface{}, options ...interface{}) *Pie {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "pie", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Pie) SetGlobalConfig(options ...interface{}) *Pie {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Pie) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Pie) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", w...)
}
