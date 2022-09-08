package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Chart3D is a chart in 3D coordinates.
type Chart3D struct {
	BaseConfiguration
}

// WithXAxis3DOpts sets the X axis of the Chart3D instance.
func WithXAxis3DOpts(opt opts.XAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.XAxis3D = opt
	}
}

// WithYAxis3DOpts sets the Y axis of the Chart3D instance.
func WithYAxis3DOpts(opt opts.YAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.YAxis3D = opt
	}
}

// WithZAxis3DOpts sets the Z axis of the Chart3D instance.
func WithZAxis3DOpts(opt opts.ZAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ZAxis3D = opt
	}
}

// WithGrid3DOpts sets the grid of the Chart3D instance.
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
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
}

// SetGlobalOptions sets options for the Chart3D instance.
func (c *Chart3D) SetGlobalOptions(options ...GlobalOpts) *Chart3D {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Chart3D) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
