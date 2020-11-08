package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Chart3D is a chart in 3D coordinates.
type Chart3D struct {
	BaseConfiguration
}

// WithXAxis3DOpts
func WithXAxis3DOpts(opt opts.XAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.XAxis3D = opt
	}
}

// WithYAxis3DOpts
func WithYAxis3DOpts(opt opts.YAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.YAxis3D = opt
	}
}

// WithZAxis3DOpts
func WithZAxis3DOpts(opt opts.ZAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ZAxis3D = opt
	}
}

// WithGrid3DOpts
func WithGrid3DOpts(opt opts.Grid3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Grid3D = opt
	}
}

func (c *Chart3D) initChart3D() {
	c.JSAssets.Add("echarts-gl.min.js")
	c.has3DAxis = true
}

func (c *Chart3D) addSeries(chartType, name string, data []opts.Chart3DData, options ...SeriesOpts) {
	series := SingleSeries{
		Name:        name,
		Type:        chartType,
		Data:        data,
		CoordSystem: types.ChartCartesian3D,
	}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
}

// SetGlobalOptions sets options for the Chart3D instance.
func (c *Chart3D) SetGlobalOptions(options ...GlobalOpts) *Chart3D {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate
func (c *Chart3D) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
