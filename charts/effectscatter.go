package charts

// EffectScatter represents an effect scatter chart.
type EffectScatter struct {
	RectChart
}

func (EffectScatter) Type() string { return ChartType.EffectScatter }

// EffectScatterChartOpts is the option set for an effect scatter chart.
type EffectScatterChartOpts struct {
	XAxisIndex int
	YAxisIndex int
}

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
func (c *EffectScatter) AddYAxis(name string, yAxis interface{}, options ...SeriesOptser) *EffectScatter {
	series := singleSeries{Name: name, Type: ChartType.EffectScatter, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
