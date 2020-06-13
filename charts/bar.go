package charts

import (
	"io"
)

// Bar represents a bar chart.
type Bar struct {
	RectChart
	//BarOpts

	isXYReversal bool
}

func (Bar) Type() string { return ChartType.Bar }

// NewBar creates a new bar chart.
func NewBar(routers ...RouterOpts) *Bar {
	chart := new(Bar)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Bar) AddXAxis(xAxis interface{}) *Bar {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Bar) AddYAxis(name string, yAxis interface{}, fns ...SeriesOptFn) *Bar {
	series := SingleSeries{Name: name, Type: ChartType.Bar, Data: yAxis}
	series.configureSeriesFns(fns...)
	//series.setSingleSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	//c.setColor(options...)
	return c
}

// XYReversal checks if X axis and Y axis are reversed.
func (c *Bar) XYReversal() *Bar {
	c.isXYReversal = true
	return c
}

func (c *Bar) validateOpts() {
	c.XAxisOptsList[0].Data = c.xAxisData
	// XY 轴翻转
	if c.isXYReversal {
		c.YAxisOptsList[0].Data = c.xAxisData
		c.XAxisOptsList[0].Data = nil
	}
	// 确保 Y 轴数标签正确显示
	for i := 0; i < len(c.YAxisOptsList); i++ {
		c.YAxisOptsList[i].AxisLabel = true
	}
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Bar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
