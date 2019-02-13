package charts

import "github.com/chenjiandongx/go-echarts/common"

type BoxPlot struct {
	RectChart
}

func (BoxPlot) chartType() string { return common.ChartType.BoxPlot }

func NewBoxPlot(routers ...RouterOpts) *BoxPlot {
	chart := new(BoxPlot)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

func (c *BoxPlot) AddXAxis(xAxis interface{}) *BoxPlot {
	c.xAxisData = xAxis
	return c
}

func (c *BoxPlot) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *BoxPlot {
	series := singleSeries{Name: name, Type: common.ChartType.BoxPlot, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
