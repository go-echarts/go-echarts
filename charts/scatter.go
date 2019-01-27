package charts

type Scatter struct {
	RectChart
}

func NewScatter(routers ...HTTPRouter) *Scatter {
	chart := new(Scatter)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	c.xAxisData = xAxis
	return c
}

func (c *Scatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Scatter {
	series := singleSeries{Name: name, Type: "scatter", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
