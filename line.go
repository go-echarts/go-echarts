package goecharts

type Line struct {
	RectChart
}

// 工厂函数，生成 `Line` 实例
func NewLine(routers ...HTTPRouter) *Line {
	chart := new(Line)
	chart.HasXYAxis = true
	chart.initBaseOpts(routers...)
	chart.initAssetsOpts()
	return chart
}

// 提供 X 轴数据
func (c *Line) AddXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *Line) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Line {
	series := singleSeries{Name: name, Type: "line", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
