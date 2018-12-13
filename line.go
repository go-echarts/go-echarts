package goecharts

type Line struct {
	RectChart
}

//工厂函数，生成 `Line` 实例
func NewLine() *Line {
	line := new(Line)
	line.HasXYAxis = true
	return line
}

// 提供 X 轴数据
func (line *Line) AddXAxis(xAxis interface{}) *Line {
	line.xAxisData = xAxis
	return line
}

// 提供 Y 轴数据
func (line *Line) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Line {
	series := Series{Name: name, Type: lineType, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	line.SeriesList = append(line.SeriesList, series)
	return line
}
