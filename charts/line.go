package charts

// Line represents a line chart.
type Line struct {
	RectChart
}

func (Line) Type() string { return ChartType.Line }

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
func (c *Line) AddYAxis(name string, yAxis interface{}, options ...SeriesOptser) *Line {
	series := singleSeries{Name: name, Type: ChartType.Line, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
