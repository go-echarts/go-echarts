package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// BoxPlot represents a boxplot chart.
type BoxPlot struct {
	RectChart
}

// Type returns the chart type.
func (BoxPlot) Type() string { return types.ChartBoxPlot }

// NewBoxPlot creates a new boxplot chart.
func NewBoxPlot() *BoxPlot {
	c := &BoxPlot{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *BoxPlot) SetXAxis(x interface{}) *BoxPlot {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *BoxPlot) AddSeries(name string, data []opts.BoxPlotData, options ...SeriesOpts) *BoxPlot {
	series := SingleSeries{Name: name, Type: types.ChartBoxPlot, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *BoxPlot) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
