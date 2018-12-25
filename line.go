package goecharts

type Line struct {
	RectChart
}

//工厂函数，生成 `Line` 实例
func NewLine(routers ...HttpRouter) *Line {
	line := new(Line)
	line.HasXYAxis = true
	line.initBaseOpts(routers...)
	line.initAssetsOpts()
	return line
}

// 提供 X 轴数据
func (line *Line) AddXAxis(xAxis interface{}) *Line {
	line.xAxisData = xAxis
	return line
}

// 提供 Y 轴数据及 Series 配置项
func (line *Line) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Line {
	series := singleSeries{Name: name, Type: lineType, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	line.Series = append(line.Series, series)
	line.setColor(options...)
	return line
}
