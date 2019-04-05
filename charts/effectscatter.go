package charts

// EffectScatter represents an effect scatter chart.
type EffectScatter struct {
	RectChart
}

func (EffectScatter) chartType() string { return ChartType.EffectScatter }

// EffectScatterChartOpts is the option set for an effect scatter chart.
type EffectScatterChartOpts struct {
	XAxisIndex int
	YAxisIndex int
}

// RippleEffectOpts is the option set for the ripple effect.
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

func (RippleEffectOpts) markSeries() {}

// NewEffectScatter creates a new effect scatter chart.
func NewEffectScatter(routers ...RouterOpts) *EffectScatter {
	chart := new(EffectScatter)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *EffectScatter) AddXAxis(xAxis interface{}) *EffectScatter {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *EffectScatter {
	series := singleSeries{Name: name, Type: ChartType.EffectScatter, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
