package goecharts

import (
	"io"
)

type Bar struct {
	RectChart
	BarChartOpts

	isXYReversal bool
}

// Bar series options
type BarChartOpts struct {
	Stack      string
	XAxisIndex int
	YAxisIndex int
}

// 工厂函数，生成 `Bar` 实例
func NewBar(routers ...HTTPRouter) *Bar {
	chart := new(Bar)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

// 提供 X 轴数据
func (c *Bar) AddXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *Bar) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Bar {
	series := singleSeries{Name: name, Type: "bar", Data: yAxis}
	series.setSingleSeriesOpts(options...)

	if ok, opt := switchChartOpts(options...); ok {
		opt := opt.(BarChartOpts)
		series.Stack = opt.Stack
		series.XAxisIndex = opt.XAxisIndex
		series.YAxisIndex = opt.YAxisIndex
	}

	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Bar) XYReversal() {
	c.isXYReversal = true
}

// 对图形配置做最后的验证，确保能够正确渲染
func (c *Bar) validateOpts() {
	c.XAxisOptsList[0].Data = c.xAxisData
	// XY 轴翻转
	if c.isXYReversal {
		c.YAxisOptsList[0].Data = c.xAxisData
		c.XAxisOptsList[0].Data = nil
	}
	c.validateAssets(c.AssetsHost)
}

func (c *Bar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", w...)
}
