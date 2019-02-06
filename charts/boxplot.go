package charts

import "github.com/chenjiandongx/go-echarts/common"

type BoxPlot struct {
	RectChart
}

func (BoxPlot) chartType() string { return common.BoxPlotType }

func NewBoxPlot(routers ...HTTPRouter) *BoxPlot {
	chart := new(BoxPlot)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *BoxPlot) AddXAxis(xAxis interface{}) *BoxPlot {
	c.xAxisData = xAxis
	return c
}

func (c *BoxPlot) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *BoxPlot {
	series := singleSeries{Name: name, Type: "boxplot", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
