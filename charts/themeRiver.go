package charts

import (
	"io"
)

// ThemeRiver represents a theme river chart.
type ThemeRiver struct {
	BaseOpts
	MultiSeries
}


func (SingleAxisOpts) MarkGlobal() {}

func (ThemeRiver) Type() string { return ChartType.ThemeRiver }

// NewThemeRiver creates a new theme river chart.
func NewThemeRiver(routers ...RouterOpts) *ThemeRiver {
	chart := new(ThemeRiver)
	chart.initBaseOpts(routers...)
	chart.HasSingleAxis = true
	return chart
}

// Add adds new data sets.
func (c *ThemeRiver) Add(name string, data interface{}, fns ...SeriesOpts) *ThemeRiver {
	series := SingleSeries{Name: name, Type: ChartType.ThemeRiver, Data: data}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the ThemeRiver instance.
func (c *ThemeRiver) SetGlobalOptions(options ...GlobalOptser) *ThemeRiver {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *ThemeRiver) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *ThemeRiver) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
