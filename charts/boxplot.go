package charts

// BoxPlot represents a boxplot chart.
type BoxPlot struct {
	RectChart
}

func (BoxPlot) chartType() string { return ChartType.BoxPlot }

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
func (c *BoxPlot) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *BoxPlot {
	series := singleSeries{Name: name, Type: ChartType.BoxPlot, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
