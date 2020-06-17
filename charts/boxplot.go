package charts

// BoxPlot represents a boxplot chart.
type BoxPlot struct {
	RectChart
}

func (BoxPlot) Type() string { return ChartType.BoxPlot }

// NewBoxPlot creates a new boxplot chart.
func NewBoxPlot(routers ...RouterOpts) *BoxPlot {
	chart := new(BoxPlot)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *BoxPlot) AddXAxis(xAxis interface{}) *BoxPlot {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *BoxPlot) AddYAxis(name string, yAxis interface{}, fns ...SeriesOpts) *BoxPlot {
	series := SingleSeries{Name: name, Type: ChartType.BoxPlot, Data: yAxis}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.setColor(options...)
	return c
}
