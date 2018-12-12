package goecharts

type Scatter struct {
	RectChart
}

//工厂函数，生成 `Scatter` 实例
func NewScatter() *Scatter {
	scatter := new(Scatter)
	scatter.HasXYAxis = true
	return scatter
}

func (scatter *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	scatter.xAxisData = xAxis
	return scatter
}

func (scatter *Scatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Scatter {
	series := Series{Name: name, Type: scatterType, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	scatter.SeriesList = append(scatter.SeriesList, series)
	return scatter
}
