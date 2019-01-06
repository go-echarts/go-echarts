package goecharts

type Scatter struct {
	RectChart
}

// 工厂函数，生成 `Scatter` 实例
func NewScatter(routers ...HTTPRouter) *Scatter {
	chart := new(Scatter)
	chart.initBaseOpts(true, routers...)
	return chart
}

// 提供 X 轴数据
func (c *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *Scatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Scatter {
	series := singleSeries{Name: name, Type: "scatter", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
