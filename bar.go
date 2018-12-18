package goecharts

import (
	"bytes"
	"io"
)

type Bar struct {
	// 是否翻转 XY 轴
	IsXYReversal bool
	RectChart
}

//工厂函数，生成 `Bar` 实例
func NewBar(routers ...HttpRouter) *Bar {
	bar := new(Bar)
	bar.HasXYAxis = true
	bar.init(routers...)
	return bar
}

// 提供 X 轴数据
func (bar *Bar) AddXAxis(xAxis interface{}) *Bar {
	bar.xAxisData = xAxis
	return bar
}

// 提供 Y 轴数据
func (bar *Bar) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Bar {
	series := Series{Name: name, Type: barType, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	bar.SeriesList = append(bar.SeriesList, series)
	bar.setColor(options...)
	return bar
}

// 对图形配置做最后的验证，确保能够正确渲染
func (bar *Bar) verifyOpts() {
	bar.XAxisOptions.Data = bar.xAxisData
	// XY 轴翻转
	if bar.IsXYReversal {
		bar.YAxisOptions.Data = bar.xAxisData
		bar.XAxisOptions.Data = nil
	}
	bar.verifyInitOpt()
}

// 渲染图表，支持多 io.Writer
func (bar *Bar) Render(w ...io.Writer) {
	bar.XAxisOptions.Data = bar.xAxisData
	// XY 轴翻转
	if bar.IsXYReversal {
		bar.YAxisOptions.Data = bar.xAxisData
		bar.XAxisOptions.Data = nil
	}

	bar.insertSeriesColors(bar.appendColor)
	bar.verifyOpts()

	var b bytes.Buffer
	renderChart(bar, &b, "chart")
	res := replaceRender(b)
	// 只渲染第一次
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}
