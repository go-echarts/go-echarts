package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
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
	chart := &Bar{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// SetXAxis sets the X axis.
func (c *Bar) SetXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

// AddSeries adds the new series.
func (c *Bar) AddSeries(name string, data []opts.BarChartItem, opts ...SeriesOpts) *Bar {
	series := SingleSeries{Name: name, Type: types.ChartBar, Data: data}
	series.configureSeriesOpts(opts...)
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

// Render renders the chart and writes the output to given writers.
func (c *Bar) Render(w io.Writer) error {
	c.Validate()
	return renderToWriter(c, "chart", []string{}, w)
}
