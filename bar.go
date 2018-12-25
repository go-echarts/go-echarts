package goecharts

import (
	"bytes"
	"io"
)

type Bar struct {
	// 是否翻转 XY 轴
	IsXYReversal bool
	RectChart

	HasXYAxis bool
}

// 工厂函数，生成 `Bar` 实例
func NewBar(routers ...HttpRouter) *Bar {
	bar := new(Bar)
	bar.HasXYAxis = true
	bar.init(routers...)
	bar.initAssetsOpts()
	return bar
}

// 提供 X 轴数据
func (bar *Bar) AddXAxis(xAxis interface{}) *Bar {
	bar.xAxisData = xAxis
	return bar
}

// 提供 Y 轴数据
func (bar *Bar) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Bar {
	series := singleSeries{Name: name, Type: barType, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	bar.Series = append(bar.Series, series)
	bar.setColor(options...)
	return bar
}

// 对图形配置做最后的验证，确保能够正确渲染
func (bar *Bar) validateOpts() {
	bar.XAxisOpts.Data = bar.xAxisData
	// XY 轴翻转
	if bar.IsXYReversal {
		bar.YAxisOpts.Data = bar.xAxisData
		bar.XAxisOpts.Data = nil
	}
	bar.validateInitOpt()
	bar.validateAssets(bar.AssetsHost)
}

// 渲染图表，支持多 io.Writer
func (bar *Bar) Render(w ...io.Writer) {
	bar.XAxisOpts.Data = bar.xAxisData
	// XY 轴翻转
	if bar.IsXYReversal {
		bar.YAxisOpts.Data = bar.xAxisData
		bar.XAxisOpts.Data = nil
	}

	bar.insertSeriesColors(bar.appendColor)
	bar.validateOpts()

	var b bytes.Buffer
	renderChart(bar, &b, "chart")
	res := replaceRender(b)
	// 只渲染第一次
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}
