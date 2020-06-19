package charts

// HeatMap represents a heatmap chart.
type HeatMap struct {
	RectChart
}

func (HeatMap) Type() string { return ChartType.HeatMap }

// NewHeatMap creates a new heatmap chart.
func NewHeatMap(routers ...RouterOpts) *HeatMap {
	chart := new(HeatMap)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *HeatMap) AddXAxis(xAxis interface{}) *HeatMap {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *HeatMap) AddYAxis(name string, yAxis interface{}, options ...SeriesOptser) *HeatMap {
	series := singleSeries{Name: name, Type: ChartType.HeatMap, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
