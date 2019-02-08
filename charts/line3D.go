package charts

import (
	"github.com/chenjiandongx/go-echarts/common"
)

type Line3D struct {
	Chart3D
}

func (Line3D) chartType() string { return common.ChartType.Line3D }

func NewLine3D(routers ...RouterOpts) *Line3D {
	chart := new(Line3D)
	chart.initBaseOpts(false, routers...)
	chart.initChart3D()
	return chart
}

func (c *Line3D) AddXYAxis(xAxis, yAxis interface{}) *Line3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

func (c *Line3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Line3D {
	c.addZAxis(common.ChartType.Line3D, name, zAxis, options...)
	return c
}
