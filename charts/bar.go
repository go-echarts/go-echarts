package charts

import (
	"io"
)

// Bar represents a bar chart.
type Bar struct {
	RectChart
	BarOpts

	isXYReversal bool
}

func (Bar) chartType() string { return ChartType.Bar }

// BarOpts is the option set for a bar chart.
type BarOpts struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 不同系列的柱间距离，为百分比（如 "30%"，表示柱子宽度的 30%）。
	// 如果想要两个系列的柱子重叠，可以设置 barGap 为 "-100%"。这在用柱子做背景的时候有用。
	// 在同一坐标系上，此属性会被多个 "bar" 系列共享。
	// 此属性应设置于此坐标系中最后一个 "bar" 系列上才会生效，并且是对此坐标系中所有 "bar" 系列生效
	// 默认 "30%"
	BarGap string
	// 同一系列的柱间距离，默认为类目间距的 20%，可设固定值
	// 在同一坐标系上，此属性会被多个 "bar" 系列共享。
	// 此属性应设置于此坐标系中最后一个 "bar" 系列上才会生效，并且是对此坐标系中所有 "bar" 系列生效
	BarCategoryGap string
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (BarOpts) markSeries() {}

func (opt *BarOpts) setChartOpt(s *singleSeries) {
	s.Stack = opt.Stack
	s.BarGap = opt.BarGap
	s.BarCategoryGap = opt.BarCategoryGap
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

// NewBar creates a new bar chart.
func NewBar(routers ...RouterOpts) *Bar {
	chart := new(Bar)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Bar) AddXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Bar) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Bar {
	series := singleSeries{Name: name, Type: ChartType.Bar, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

// XYReversal checks if X axis and Y axis are reversed.
func (c *Bar) XYReversal() *Bar {
	c.isXYReversal = true
	return c
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

// Render renders the chart and writes the output to given writers.
func (c *Bar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
