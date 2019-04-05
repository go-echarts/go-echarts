package charts

// Scatter3D represents a 3D scatter chart.
type Scatter3D struct {
	Chart3D
}

func (Scatter3D) chartType() string { return ChartType.Scatter3D }

// NewScatter3D creates a new 3D scatter chart.
func NewScatter3D(routers ...RouterOpts) *Scatter3D {
	chart := new(Scatter3D)
	chart.initBaseOpts(routers...)
	chart.initChart3D()
	return chart
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Scatter3D) AddXYAxis(xAxis, yAxis interface{}) *Scatter3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Scatter3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Scatter3D {
	c.addZAxis(ChartType.Scatter3D, name, zAxis, options...)
	return c
}
