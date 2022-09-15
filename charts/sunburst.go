package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Sunburst represents a sunburst chart.
type Sunburst struct {
	BaseConfiguration
	BaseActions
}

// Type returns the chart type.
func (Sunburst) Type() string { return types.ChartSunburst }

// NewSunburst creates a new sunburst chart instance.
func NewSunburst() *Sunburst {
	c := &Sunburst{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *Sunburst) AddSeries(name string, data []opts.SunBurstData, options ...SeriesOpts) *Sunburst {
	series := SingleSeries{Name: name, Type: types.ChartSunburst, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Sunburst) SetGlobalOptions(options ...GlobalOpts) *Sunburst {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions sets actions for the Sunburst instance.
func (c *Sunburst) SetDispatchActions(actions ...GlobalActions) *Sunburst {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate validates the given configuration.
func (c *Sunburst) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
