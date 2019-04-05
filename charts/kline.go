package charts

// Kline represents a kline chart.
type Kline struct {
	RectChart
}

func (Kline) chartType() string { return ChartType.Kline }

// NewKLine creates a new kline chart.
func NewKLine(routers ...RouterOpts) *Kline {
	chart := new(Kline)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Kline) AddXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Kline) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Kline {
	series := singleSeries{Name: name, Type: ChartType.Kline, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
