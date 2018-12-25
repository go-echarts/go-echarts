package goecharts

type EffectScatter struct {
	RectChart
}

// 涟漪特效配置项
type RippleEffectOpts struct {
	// 动画的周期，秒数
	// 默认 4(s)
	Period float32 `json:"period,omitempty"`
	// 动画中波纹的最大缩放比例
	// 默认 2.5
	Scale float32 `json:"scale,omitempty"`
	// 波纹的绘制方式，可选 "stroke" 和 "fill"
	// 默认 "fill"
	BrushType string `json:"brushType,omitempty"`
}

//工厂函数，生成 `Scatter` 实例
func NewEffectScatter(routers ...HttpRouter) *EffectScatter {
	es := new(EffectScatter)
	es.HasXYAxis = true
	es.initBaseOpts(routers...)
	es.initAssetsOpts()
	return es
}

// 提供 X 轴数据
func (es *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	es.xAxisData = xAxis
	return es
}

// 提供 Y 轴数据
func (es *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *EffectScatter {
	series := singleSeries{Name: name, Type: effectScatterType, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	es.Series = append(es.Series, series)
	es.setColor(options...)
	return es
}
