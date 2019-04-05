package charts

// Line3D represents a 3D line chart.
type Line3D struct {
	Chart3D
}

func (Line3D) chartType() string { return ChartType.Line3D }

// NewLine3D creates a new 3D line chart.
func NewLine3D(routers ...RouterOpts) *Line3D {
	chart := new(Line3D)
	chart.initBaseOpts(routers...)
	chart.initChart3D()
	return chart
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Line3D) AddXYAxis(xAxis, yAxis interface{}) *Line3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Line3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Line3D {
	c.addZAxis(ChartType.Line3D, name, zAxis, options...)
	return c
}
