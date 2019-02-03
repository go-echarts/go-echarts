package charts

import "io"

type Bar struct {
	RectChart
	BarOpts

	isXYReversal bool
}

func (Bar) chartType() string { return "bar" }

// Bar series options
type BarOpts struct {
	Stack      string
	XAxisIndex int
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
	series := singleSeries{Name: name, Type: "bar", Data: yAxis}
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
