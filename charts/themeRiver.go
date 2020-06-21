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
	chart := &ThemeRiver{}
	chart.initBaseConfiguration()
	chart.HasSingleAxis = true
	return chart
}

// Add adds new data sets.
func (c *ThemeRiver) AddSeries(name string, data interface{}, opts ...SeriesOpts) *ThemeRiver {
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

func (c *ThemeRiver) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *ThemeRiver) Render(w io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.Validate()
	return renderToWriter(c, "chart", []string{}, w)
}
