package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/types"
)

// Bar represents a bar chart.
type Bar struct {
	RectChart

	isXYReversal bool
}

func (Bar) Type() string { return types.ChartBar }

// NewBar creates a new bar chart.
func NewBar() *Bar {
	chart := &Bar{}
	chart.initBaseConfiguration()
	chart.initXYAxis()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Bar) AddXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Bar) AddYAxis(name string, yAxis interface{}, opts ...SeriesOpts) *Bar {
	series := SingleSeries{Name: name, Type: types.ChartBar, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// XYReversal checks if X axis and Y axis are reversed.
func (c *Bar) XYReversal() *Bar {
	c.isXYReversal = true
	return c
}

func (c *Bar) validateOpts() {
	c.XAxisList[0].Data = c.xAxisData
	// XY 轴翻转
	if c.isXYReversal {
		c.YAxisList[0].Data = c.xAxisData
		c.XAxisList[0].Data = nil
	}
	// 确保 Y 轴数标签正确显示
	for i := 0; i < len(c.YAxisList); i++ {
		c.YAxisList[i].AxisLabel.Show = true
	}
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Bar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
