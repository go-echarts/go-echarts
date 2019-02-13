package charts

import (
	"github.com/chenjiandongx/go-echarts/common"
)

type Scatter3D struct {
	Chart3D
}

func (Scatter3D) chartType() string { return common.ChartType.Scatter3D }

func NewScatter3D(routers ...RouterOpts) *Scatter3D {
	chart := new(Scatter3D)
	chart.initBaseOpts(routers...)
	chart.initChart3D()
	return chart
}

func (c *Scatter3D) AddXYAxis(xAxis, yAxis interface{}) *Scatter3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

func (c *Scatter3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Scatter3D {
	c.addZAxis(common.ChartType.Scatter3D, name, zAxis, options...)
	return c
}
