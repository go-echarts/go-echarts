package charts

// Scatter represents a scatter chart.
type Scatter struct {
	RectChart
}

func (Scatter) Type() string { return ChartType.Scatter }

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
func (c *Scatter) AddYAxis(name string, yAxis interface{}, fns ...SeriesOptFn) *Scatter {
	series := SingleSeries{Name: name, Type: ChartType.Scatter, Data: yAxis}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.setColor(options...)
	return c
}
