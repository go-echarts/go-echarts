package geocharts

import "io"

type EffectScatter struct {
	RectChart
}

//工厂函数，生成 `Scatter` 实例
func NewEffectScatter() *EffectScatter {
	es := new(EffectScatter)
	es.setDefault()
	es.HasXYAxis = true
	return es
}

func (es *EffectScatter) setDefault() {
	es.InitOptions.SetDefault()
	es.BaseOptions.SetDefault()
}

func (es *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	es.xAxisData = xAxis
	return es
}

func (es *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *EffectScatter {
	series := Series{Name: name, Type: EFFECTSCATTER, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	es.SeriesList = append(es.SeriesList, series)
	return es
}

func (es *EffectScatter) Render(w io.Writer) {
	es.XAxisOptions.Data = es.xAxisData
	RenderChart(es, w)
}
