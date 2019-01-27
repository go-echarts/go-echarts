package charts

type Line struct {
	RectChart
}

// Line series options
type LineOpts struct {
	Stack      string
	Smooth     bool
	Step       bool
	XAxisIndex int
	YAxisIndex int
}

func (opt *LineOpts) setChartOpt(s *singleSeries) {
	s.Stack = opt.Stack
	s.Smooth = opt.Smooth
	s.Step = opt.Step
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

func NewLine(routers ...HTTPRouter) *Line {
	chart := new(Line)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *Line) AddXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

func (c *Line) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Line {
	series := singleSeries{Name: name, Type: "line", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
