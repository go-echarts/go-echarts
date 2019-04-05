package charts

// Surface3D represents a 3D surface chart.
type Surface3D struct {
	Chart3D
}

func (Surface3D) chartType() string { return ChartType.Surface3D }

// NewSurface3D creates a new 3d surface chart.
func NewSurface3D(routers ...RouterOpts) *Surface3D {
	chart := new(Surface3D)
	chart.initBaseOpts(routers...)
	chart.initChart3D()
	return chart
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Surface3D) AddXYAxis(xAxis, yAxis interface{}) *Surface3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Surface3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Surface3D {
	c.addZAxis(ChartType.Surface3D, name, zAxis, options...)
	return c
}
