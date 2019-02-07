package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Bar3D struct {
	Chart3D
}

func (Bar3D) chartType() string { return common.ChartType.Bar3D }

type Bar3DOpts struct {
	Shading string
}

func (Bar3DOpts) markSeries() {}

func (opt *Bar3DOpts) setChartOpt(s *singleSeries) {
	s.Shading = opt.Shading
}

func NewBar3D(routers ...HTTPRouter) *Bar3D {
	chart := new(Bar3D)
	chart.initBaseOpts(false, routers...)
	chart.initChart3D()
	return chart
}

func (c *Bar3D) AddXYAxis(xAxis, yAxis interface{}) *Bar3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

func (c *Bar3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Bar3D {
	series := singleSeries{
		Name:        name,
		Type:        common.ChartType.Bar3D,
		Data:        zAxis,
		CoordSystem: common.ChartType.Cartesian3D,
	}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Bar3D) validateOpts() {
	// 确保 XY 轴数据项不会被抹除
	if c.XAxis3D.Data == nil {
		c.XAxis3D.Data = c.xData
	}
	if c.YAxis3D.Data == nil {
		c.YAxis3D.Data = c.yData
	}
	c.validateAssets(c.AssetsHost)
}

func (c *Bar3D) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
