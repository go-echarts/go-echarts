package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Sunburst represents a sunburst chart.
type Sunburst struct {
	RectChart
}

// Type returns the chart type.
func (Sunburst) Type() string { return types.Sunburst }

// NewSunburst creates a new sunburst chart instance.
func NewSunburst() *Sunburst {
	c := &Sunburst{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *Sunburst) AddSeries(name string, data []opts.SunBurstData, options ...SeriesOpts) *Sunburst {
	series := SingleSeries{Name: name, Type: types.Sunburst, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *Sunburst) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
