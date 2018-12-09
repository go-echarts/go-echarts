package geocharts

import "io"

type Scatter struct {
	RectChart
}

//工厂函数，生成 `Scatter` 实例
func NewScatter() *Scatter {
	scatter := new(Scatter)
	scatter.setDefault()
	scatter.HasXYAxis = true
	return scatter
}

func (scatter *Scatter) setDefault() {
	scatter.InitOptions.SetDefault()
	scatter.BaseOptions.SetDefault()
}

func (scatter *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	scatter.xAxisData = xAxis
	return scatter
}

func (scatter *Scatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Scatter {
	series := Series{Name: name, Type: SCATTER, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	scatter.SeriesList = append(scatter.SeriesList, series)
	return scatter
}

func (scatter *Scatter) Render(w io.Writer) {
	scatter.XAxisOptions.Data = scatter.xAxisData
	RenderChart(scatter, w)
}
