package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Bar represents a bar chart.
type Bar struct {
	RectChart

	isXYReversal bool
}

// Type returns the chart type.
func (Bar) Type() string { return types.ChartBar }

// NewBar creates a new bar chart instance.
func NewBar() *Bar {
	c := &Bar{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// EnablePolarType enables the polar bar.
func (c *Bar) EnablePolarType() *Bar {
	c.hasXYAxis = false
	c.hasPolar = true
	return c
}

// SetXAxis sets the X axis.
func (c *Bar) SetXAxis(x interface{}) *Bar {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *Bar) AddSeries(name string, data []opts.BarData, options ...SeriesOpts) *Bar {
	series := SingleSeries{Name: name, Type: types.ChartBar, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// XYReversal checks if X axis and Y axis are reversed.
func (c *Bar) XYReversal() *Bar {
	c.isXYReversal = true
	return c
}

// Validate validates the given configuration.
func (c *Bar) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	if c.isXYReversal {
		c.YAxisList[0].Data = c.xAxisData
		c.XAxisList[0].Data = nil
	}
	c.Assets.Validate(c.AssetsHost)
}
