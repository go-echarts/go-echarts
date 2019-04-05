package charts

// Line represents a line chart.
type Line struct {
	RectChart
}

func (Line) chartType() string { return ChartType.Line }

// LineOpts is the options set for a line chart.
type LineOpts struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 曲线是否平滑
	Smooth bool
	// 是否使用阶梯图
	Step bool
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (LineOpts) markSeries() {}

func (opt *LineOpts) setChartOpt(s *singleSeries) {
	s.Stack = opt.Stack
	s.Smooth = opt.Smooth
	s.Step = opt.Step
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

// NewLine creates a new line chart.
func NewLine(routers ...RouterOpts) *Line {
	chart := new(Line)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Line) AddXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Line) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Line {
	series := singleSeries{Name: name, Type: ChartType.Line, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
