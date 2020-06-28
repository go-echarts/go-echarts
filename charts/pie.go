package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Pie represents a pie chart.
type Pie struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (Pie) Type() string { return types.ChartPie }

// NewPie creates a new gauge chart.
func NewPie() *Pie {
	chart := &Pie{}
	chart.initBaseConfiguration()
	chart.Renderer = render.NewChartRender(chart, chart.Validate)
	return chart
}

// AddSeries adds new data sets.
func (c *Pie) AddSeries(name string, data []opts.PieData, opts ...SeriesOpts) *Pie {
	series := SingleSeries{Name: name, Type: types.ChartPie, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Pie) SetGlobalOptions(opts ...GlobalOpts) *Pie {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

// Validate
func (c *Pie) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writer.
//func (c *Pie) Render(w io.Writer) error {
//	c.Validate()
//	return render.ChartRender(c, w)
//}
