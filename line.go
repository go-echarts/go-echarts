package goecharts

type Line struct {
	RectChart
}

// Line series options
type LineChartOpts struct {
	Stack      string
	XAxisIndex int
	YAxisIndex int
}

// 工厂函数，生成 `Line` 实例
func NewLine(routers ...HTTPRouter) *Line {
	chart := new(Line)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
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

	if ok, opt := c.isSelfChartOpts(options...); ok {
		series.Stack = opt.Stack
		series.XAxisIndex = opt.XAxisIndex
		series.YAxisIndex = opt.YAxisIndex
	}

	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Line) isSelfChartOpts(options ...interface{}) (bool, LineChartOpts) {
	for i := 0; i < len(options); i++ {
		switch options[i].(type) {
		case LineChartOpts:
			return true, options[i].(LineChartOpts)
		}
	}
	return false, LineChartOpts{}
}
