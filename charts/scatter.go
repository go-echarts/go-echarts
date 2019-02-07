package charts

import "github.com/chenjiandongx/go-echarts/common"

type Scatter struct {
	RectChart
}

func (Scatter) chartType() string { return common.ChartType.Scatter }

// Scatter series options
type ScatterOpts struct {
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

func (ScatterOpts) markSeries() {}

func (opt *ScatterOpts) setChartOpt(s *singleSeries) {
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
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

func (c *Scatter) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Scatter {
	series := singleSeries{Name: name, Type: common.ChartType.Scatter, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
