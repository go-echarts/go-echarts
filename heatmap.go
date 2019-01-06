package goecharts

type HeatMap struct {
	RectChart
}

// heatMap series options
type heatMapChartOpts struct {
	XAxisIndex int
	YAxisIndex int
}

func (opt *heatMapChartOpts) setChartOpt(s *singleSeries) {
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
}

// 工厂函数，生成 `heatMap` 实例
func NewHeatMap(routers ...HTTPRouter) *HeatMap {
	chart := new(HeatMap)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

// 提供 X 轴数据
func (c *HeatMap) AddXAxis(xAxis interface{}) *HeatMap {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *HeatMap) AddYAxis(name string, yAxis interface{}, options ...interface{}) *HeatMap {
	series := singleSeries{Name: name, Type: "heatmap", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
