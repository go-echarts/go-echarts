package geocharts

type Line struct {
	RectChart
}

//工厂函数，生成 `Line` 实例
func NewLine() *Line {
	line := new(Line)
	line.setDefault()
	line.HasXYAxis = true
	return line
}

func (line *Line) setDefault() {
	line.InitOptions.SetDefault()
	line.BaseOptions.SetDefault()
}

func (line *Line) AddXAxis(xAxis interface{}) *Line {
	line.xAxisData = xAxis
	return line
}

func (line *Line) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Line {
	series := Series{Name: name, Type: LINE, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	line.SeriesList = append(line.SeriesList, series)
	return line
}
