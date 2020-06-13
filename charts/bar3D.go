package charts

// Bar3D represents a 3D bar chart.
type Bar3D struct {
	Chart3D
}

func (Bar3D) Type() string { return ChartType.Bar3D }

// NewBar3D creates a new 3D bar chart.
func NewBar3D(routers ...RouterOpts) *Bar3D {
	chart := new(Bar3D)
	chart.initBaseOpts(routers...)
	chart.initChart3D()
	return chart
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Bar3D) AddXYAxis(xAxis, yAxis interface{}) *Bar3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Bar3D) AddZAxis(name string, zAxis interface{}, options ...SeriesOptser) *Bar3D {
	c.addZAxis(ChartType.Bar3D, name, zAxis, options...)
	return c
}
