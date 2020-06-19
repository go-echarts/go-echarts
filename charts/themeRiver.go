package charts

import (
	"github.com/go-echarts/go-echarts/types"
	"io"
)

// ThemeRiver represents a theme river chart.
type ThemeRiver struct {
	BaseConfiguration
	MultiSeries
}

func (ThemeRiver) Type() string { return types.ChartThemeRiver }

// NewThemeRiver creates a new theme river chart.
func NewThemeRiver() *ThemeRiver {
	chart := new(ThemeRiver)
	chart.initBaseConfiguration()
	chart.HasSingleAxis = true
	return chart
}

// Add adds new data sets.
func (c *ThemeRiver) Add(name string, data interface{}, opts ...SeriesOpts) *ThemeRiver {
	series := SingleSeries{Name: name, Type: types.ChartThemeRiver, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the ThemeRiver instance.
func (c *ThemeRiver) SetGlobalOptions(opts ...GlobalOpts) *ThemeRiver {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
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
