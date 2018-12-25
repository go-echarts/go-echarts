package goecharts

import (
	"io"
)

type Funnel struct {
	BaseOpts
	Series

	HasXYAxis bool
}

// 工厂函数，生成 `Funnel` 实例
func NewFunnel(routers ...HttpRouter) *Funnel {
	funnel := new(Funnel)
	funnel.HasXYAxis = false
	funnel.initBaseOpts(routers...)
	funnel.initAssetsOpts()
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

func (funnel *Funnel) validateOpts() {
	funnel.validateInitOpt()
	funnel.validateAssets(funnel.AssetsHost)
}

// 渲染图表，支持多 io.Writer
func (funnel *Funnel) Render(w ...io.Writer) error {
	funnel.insertSeriesColors(funnel.appendColor)
	funnel.validateOpts()
	if err := renderToWriter(funnel, "chart", w...); err != nil {
		return err
	}
	return nil
}
