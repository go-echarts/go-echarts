package goecharts

type EffectScatter struct {
	RectChart
}

// 涟漪特效配置项
type RippleEffectOptions struct {
	// 动画的周期，秒数
	Period float32 `json:"period,omitempty" default:"4"`
	// 动画中波纹的最大缩放比例
	Scale float32 `json:"scale,omitempty" default:"2.5"`
	// 波纹的绘制方式，可选 'stroke' 和 'fill'
	BrushType string `json:"brushType,omitempty" default:"fill"`
}

//工厂函数，生成 `Scatter` 实例
func NewEffectScatter(routers ...HttpRouter) *EffectScatter {
	es := new(EffectScatter)
	es.HasXYAxis = true
	es.init(routers...)
	return es
}

// 提供 X 轴数据
func (es *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	es.xAxisData = xAxis
	return es
}

// 提供 Y 轴数据
func (es *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *EffectScatter {
	series := Series{Name: name, Type: effectScatterType, Data: yAxis}
	series.setSingleSeriesOptions(options...)
	es.SeriesList = append(es.SeriesList, series)
	es.setColor(options...)
	return es
}
