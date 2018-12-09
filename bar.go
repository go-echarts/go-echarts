package geocharts

import (
	"io"
)

type Bar struct {
	// 是否翻转 XY 轴
	IsXYReversal bool
	RectChart
}

//工厂函数，生成 `Bar` 实例
func NewBar() *Bar {
	bar := new(Bar)
	bar.setDefault()
	bar.HasXYAxis = true
	return bar
}

func (bar *Bar) setDefault() {
	bar.InitOptions.SetDefault()
	bar.RectOptions.SetDefault()
}

func (bar *Bar) AddXAxis(xAxis interface{}) *Bar {
	bar.xAxisData = xAxis
	return bar
}

func (bar *Bar) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Bar {
	series := Series{Name: name, Type: BAR, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	bar.SeriesList = append(bar.SeriesList, series)
	return bar
}

func (bar *Bar) Render(w io.Writer) {
	bar.XAxisOptions.Data = bar.xAxisData
	// XY 轴翻转
	if bar.IsXYReversal {
		bar.YAxisOptions.Data = bar.xAxisData
		bar.XAxisOptions.Data = nil
	}
	bar.SetDefault()
	RenderChart(bar, w)
}
