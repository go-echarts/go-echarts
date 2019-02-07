package charts

import "github.com/chenjiandongx/go-echarts/common"

type Bar3D struct {
	Chart3D
}

func (Bar3D) chartType() string { return common.ChartType.LineType }

func NewBar3D(routers ...HTTPRouter) *Bar3D {
	chart := new(Bar3D)
	chart.initBaseOpts(true, routers...)
	chart.Has3DAxis = true
	return chart
}

func (c *Bar3D) AddXYAxis(xAxis interface{}) *Bar3D {
	return c
}

func (c *Bar3D) AddZAxis(name string, yAxis interface{}, options ...seriesOptser) *Bar3D {
	series := singleSeries{Name: name, Type: "line", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
