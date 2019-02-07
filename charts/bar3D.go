package charts

import (
	"github.com/chenjiandongx/go-echarts/common"
)

type Bar3D struct {
	Chart3D
}

func (Bar3D) chartType() string { return common.ChartType.Bar3D }

func NewBar3D(routers ...HTTPRouter) *Bar3D {
	chart := new(Bar3D)
	chart.initBaseOpts(true, routers...)
	chart.initChart3D()
	return chart
}

func (c *Bar3D) AddXYAxis(xAxis XAxis3DOpts, yAxis YAxisOpts) *Bar3D {
	return c
}

func (c *Bar3D) AddZAxis(name string, yAxis interface{}, options ...seriesOptser) *Bar3D {
	series := singleSeries{Name: name, Type: common.ChartType.Bar3D, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
