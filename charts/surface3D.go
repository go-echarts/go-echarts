package charts

import (
	"github.com/chenjiandongx/go-echarts/common"
)

type Surface3D struct {
	Chart3D
}

func (Surface3D) chartType() string { return common.ChartType.Surface3D }

func NewSurface3D(routers ...RouterOpts) *Surface3D {
	chart := new(Surface3D)
	chart.initBaseOpts(false, routers...)
	chart.initChart3D()
	return chart
}

func (c *Surface3D) AddXYAxis(xAxis, yAxis interface{}) *Surface3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

func (c *Surface3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Surface3D {
	c.addZAxis(common.ChartType.Surface3D, name, zAxis, options...)
	return c
}
