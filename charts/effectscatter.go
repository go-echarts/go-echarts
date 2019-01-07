package charts

type EffectScatter struct {
	RectChart
}

// EffectScatter series options
type EffectScatterChartOpts struct {
	XAxisIndex int
	YAxisIndex int
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

// 工厂函数，生成 `Scatter` 实例
func NewEffectScatter(routers ...HTTPRouter) *EffectScatter {
	chart := new(EffectScatter)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

// 提供 X 轴数据
func (c *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *EffectScatter {
	series := singleSeries{Name: name, Type: "effectScatter", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
