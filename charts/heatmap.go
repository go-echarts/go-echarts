package charts

// HeatMap represents a heatmap chart.
type HeatMap struct {
	RectChart
}

func (HeatMap) chartType() string { return ChartType.HeatMap }

// HeatMapOpts is the option set for a heatmap chart.
type HeatMapOpts struct {
	//使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	//使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (HeatMapOpts) markSeries() {}

func (opt *HeatMapOpts) setChartOpt(s *singleSeries) {
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

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
func (c *HeatMap) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *HeatMap {
	series := singleSeries{Name: name, Type: ChartType.HeatMap, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
