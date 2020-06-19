package charts

import (
	"io"
)

// Radar represents a radar chart.
type Radar struct {
	BaseOpts
	MultiSeries
}

func (Radar) Type() string { return ChartType.Radar }

// NewRadar creates a new radar chart.
func NewRadar(routers ...RouterOpts) *Radar {
	chart := new(Radar)
	chart.initBaseOpts(routers...)
	chart.HasRadar = true
	return chart
}

// Add adds new data sets.
func (c *Radar) Add(name string, data interface{}, fns ...SeriesOpts) *Radar {
	series := SingleSeries{Name: name, Type: ChartType.Radar, Data: data}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.setColor(options...)
	c.legends = append(c.legends, name)
	return c
}

// SetGlobalOptions sets options for the Radar instance.
func (c *Radar) SetGlobalOptions(options ...GlobalOptser) *Radar {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Radar) validateOpts() {
	c.LegendOpts.Data = c.legends
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Radar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
