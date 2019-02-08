package charts

import "github.com/chenjiandongx/go-echarts/common"

type Kline struct {
	RectChart
}

func (Kline) chartType() string { return common.ChartType.Kline }

func NewKLine(routers ...RouterOpts) *Kline {
	chart := new(Kline)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *Kline) AddXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

func (c *Kline) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Kline {
	series := singleSeries{Name: name, Type: common.ChartType.Kline, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
