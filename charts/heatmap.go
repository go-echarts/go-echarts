package charts

type HeatMap struct {
	RectChart
}

func (HeatMap) chartType() string { return "heatMap" }

// heatMap series options
type HeatMapOpts struct {
	XAxisIndex int
	YAxisIndex int
}

func (HeatMapOpts) markSeries() {}

func (opt *HeatMapOpts) setChartOpt(s *singleSeries) {
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

func NewHeatMap(routers ...HTTPRouter) *HeatMap {
	chart := new(HeatMap)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *HeatMap) AddXAxis(xAxis interface{}) *HeatMap {
	c.xAxisData = xAxis
	return c
}

func (c *HeatMap) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *HeatMap {
	series := singleSeries{Name: name, Type: "heatmap", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
