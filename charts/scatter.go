package charts

// Scatter represents a scatter chart.
type Scatter struct {
	RectChart
}

func (Scatter) chartType() string { return ChartType.Scatter }

// ScatterOpts is the option set for a scatter chart.
type ScatterOpts struct {
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (ScatterOpts) markSeries() {}

func (opt *ScatterOpts) setChartOpt(s *singleSeries) {
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

// NewScatter creates a new scatter chart.
func NewScatter(routers ...RouterOpts) *Scatter {
	chart := new(Scatter)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Scatter) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Scatter {
	series := singleSeries{Name: name, Type: ChartType.Scatter, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
