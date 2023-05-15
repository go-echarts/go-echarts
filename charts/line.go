package charts

import (
	"github.com/go-echarts/go-echarts/v2/config"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/series"
)

type LineConfiguration struct {
	*config.BaseConfiguration
	Series          *series.LineSeries `json:"series,omitempty"`
	XAxis           *opts.XAxis        `json:"xAxis,omitempty"`
	YAxis           *opts.YAxis        `json:"yAxis,omitempty"`
	render.Renderer `json:"-"`
}

// Line represents a line chart.
type Line struct {
	*LineConfiguration
}

func NewLine() *Line {
	line := &Line{
		&LineConfiguration{
			BaseConfiguration: config.BaseConfiguration{}.New(),
			XAxis:             opts.XAxis{}.New(),
			YAxis:             opts.YAxis{}.New(),
			Series:            series.LineSeries{Type: "line"}.New(),
		},
	}

	// TODO need refine
	line.Initialization = &opts.Initialization{}
	opts.SetDefaultValue(line.Initialization)
	line.Initialization.ChartID = opts.GenUUID()
	line.Assets = &opts.Assets{}
	line.Assets.JSAssets.Init("https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js")

	line.Renderer = render.NewChartRender(line)
	return line
}
