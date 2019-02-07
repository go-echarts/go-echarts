package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Bar struct {
	RectChart
	BarOpts

	isXYReversal bool
}

func (Bar) chartType() string { return common.ChartType.Bar }

// Bar series options
type BarOpts struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (BarOpts) markSeries() {}

func (opt *BarOpts) setChartOpt(s *singleSeries) {
	s.Stack = opt.Stack
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

func NewBar(routers ...HTTPRouter) *Bar {
	chart := new(Bar)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *Bar) AddXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

func (c *Bar) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Bar {
	series := singleSeries{Name: name, Type: common.ChartType.Bar, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Bar) XYReversal() {
	c.isXYReversal = true
}

func (c *Bar) validateOpts() {
	c.XAxisOptsList[0].Data = c.xAxisData
	// XY 轴翻转
	if c.isXYReversal {
		c.YAxisOptsList[0].Data = c.xAxisData
		c.XAxisOptsList[0].Data = nil
	}
	// 确保 Y 轴数标签正确显示
	for i := 0; i < len(c.YAxisOptsList); i++ {
		c.YAxisOptsList[i].AxisLabel.Show = true
	}
	c.validateAssets(c.AssetsHost)
}

func (c *Bar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
