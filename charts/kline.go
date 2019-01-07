package charts

type Kline struct {
	RectChart
}

// 工厂函数，生成 `Kline` 实例
func NewKLine(routers ...HTTPRouter) *Kline {
	chart := new(Kline)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

// 提供 X 轴数据
func (c *Kline) AddXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *Kline) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Kline {
	series := singleSeries{Name: name, Type: "candlestick", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
