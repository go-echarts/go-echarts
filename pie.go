package goecharts

import (
	"io"
)

type Pie struct {
	BaseOpts
	Series

	HasXYAxis bool
}

//工厂函数，生成 `Pie` 实例
func NewPie(routers ...HttpRouter) *Pie {
	pie := new(Pie)
	pie.HasXYAxis = false
	pie.init(routers...)
	pie.initAssetsOpts()
	return pie
}

func (pie *Pie) Add(name string, data map[string]interface{}, options ...interface{}) *Pie {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: pieType, Data: nvs}
	series.setSingleSeriesOpts(options...)
	pie.Series = append(pie.Series, series)
	pie.setColor(options...)
	return pie
}

func (pie *Pie) SetGlobalConfig(options ...interface{}) *Pie {
	pie.BaseOpts.setBaseGlobalConfig(options...)
	return pie
}

func (pie *Pie) verifyOpts() {
	pie.verifyInitOpt()
	pie.verifyAssets(pie.AssetsHost)
}

// 渲染图表，支持多 io.Writer
func (pie *Pie) Render(w ...io.Writer) {
	pie.insertSeriesColors(pie.appendColor)
	pie.verifyOpts()
	renderToWriter(pie, "chart", w...)
}
