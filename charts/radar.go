package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Radar represents a radar chart.
type Radar struct {
	BaseConfiguration
	MultiSeries
}

func (Radar) Type() string { return types.ChartRadar }

// NewRadar creates a new radar chart.
func NewRadar() *Radar {
	chart := &Radar{}
	chart.initBaseConfiguration()
	chart.HasRadar = true
	return chart
}

// Add adds new data sets.
func (c *Radar) AddSeries(name string, data []opts.RadarChartItem, opts ...SeriesOpts) *Radar {
	series := SingleSeries{Name: name, Type: types.ChartRadar, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.legends = append(c.legends, name)
	return c
}

// SetGlobalOptions sets options for the Radar instance.
func (c *Radar) SetGlobalOptions(opts ...GlobalOpts) *Radar {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Radar) Validate() {
	c.Legend.Data = c.legends
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Radar) Render(w io.Writer) error {
	c.Validate()
	return render.ChartRender(c, w)
}
