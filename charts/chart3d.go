package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/render"

	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

// Chart3D is a chart in 3D coordinate.
type Chart3D struct {
	BaseConfiguration
	MultiSeries

	XAxis3D opts.XAxis3D
	YAxis3D opts.YAxis3D
	ZAxis3D opts.ZAxis3D
	Grid3D  opts.Grid3D

	xData interface{}
	yData interface{}
}

func WithXAxis3DOpts(opt opts.XAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.XAxis3D = opt
	}
}

func WithYAxis3DOpts(opt opts.YAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.YAxis3D = opt
	}
}

func WithZAxis3DOpts(opt opts.ZAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ZAxis3D = opt
	}
}

func WithGrid3DOpts(opt opts.Grid3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Grid3D = opt
	}
}

func (c *Chart3D) initChart3D() {
	c.JSAssets.Add("echarts-gl.min.js")
	c.Has3DAxis = true
}

// SetGlobalOptions sets options for the Chart3D instance.
func (c *Chart3D) SetGlobalOptions(opts ...GlobalOpts) *Chart3D {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Chart3D) Validate() {
	// 确保 XY 轴数据项不会被抹除
	if c.XAxis3D.Data == nil {
		c.XAxis3D.Data = c.xData
	}
	if c.YAxis3D.Data == nil {
		c.YAxis3D.Data = c.yData
	}
	c.Assets.Validate(c.AssetsHost)
}

func (c *Chart3D) addZAxis(chartType, name string, zAxis interface{}, opts ...SeriesOpts) {
	series := SingleSeries{
		Name:        name,
		Type:        chartType,
		Data:        zAxis,
		CoordSystem: types.ChartCartesian3D,
	}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
}

// Render renders the chart and writes the output to given writers.
func (c *Chart3D) Render(w io.Writer) error {
	c.Validate()
	return render.ChartRender(c, w)
}
