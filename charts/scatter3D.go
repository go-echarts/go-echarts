package charts

import "github.com/go-echarts/go-echarts/types"

// Scatter3D represents a 3D scatter chart.
type Scatter3D struct {
	Chart3D
}

func (Scatter3D) Type() string { return types.ChartScatter3D }

// NewScatter3D creates a new 3D scatter chart.
func NewScatter3D() *Scatter3D {
	chart := &Scatter3D{}
	chart.initBaseConfiguration()
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
func (c *Scatter3D) AddZAxis(name string, zAxis interface{}, opts ...SeriesOpts) *Scatter3D {
	c.addZAxis(types.ChartScatter3D, name, zAxis, opts...)
	return c
}
