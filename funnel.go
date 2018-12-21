package goecharts

import (
	"bytes"
	"io"
)

type Funnel struct {
	InitOpts
	BaseOpts
	Series

	HasXYAxis bool
}

// 工厂函数，生成 `Funnel` 实例
func NewFunnel(routers ...HttpRouter) *Funnel {
	funnel := new(Funnel)
	funnel.HasXYAxis = false
	funnel.init(routers...)
	return funnel
}

func (funnel *Funnel) Add(name string, data map[string]interface{}, options ...interface{}) *Funnel {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: funnelType, Data: nvs}
	series.setSingleSeriesOpts(options...)
	funnel.Series = append(funnel.Series, series)
	funnel.setColor(options...)
	return funnel
}

func (funnel *Funnel) SetGlobalConfig(options ...interface{}) *Funnel {
	funnel.BaseOpts.setBaseGlobalConfig(options...)
	return funnel
}

func (funnel *Funnel) verifyOpts() {
	funnel.verifyInitOpt()
}

// 渲染图表，支持多 io.Writer
func (funnel *Funnel) Render(w ...io.Writer) {

	funnel.insertSeriesColors(funnel.appendColor)
	funnel.verifyOpts()

	var b bytes.Buffer
	renderChart(funnel, &b, "chart")
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}
