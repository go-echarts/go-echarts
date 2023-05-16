package charts

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/config"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineConfiguration struct {
	*config.BaseConfiguration
	Series *series.LineSeries `json:"series,omitempty"`
	XAxis  *opts.XAxis        `json:"xAxis,omitempty"`
	YAxis  *opts.YAxis        `json:"yAxis,omitempty"`
}

// Line represents a line chart.
type Line struct {
	*LineConfiguration
}

func (line *Line) GetChartName() string {
	return types.ChartLine
}

func (line *Line) GetChart() interface{} {
	return line
}

func (line *Line) GetContainer() *components.Container {
	if line.Container != nil {
		return line.Container
	}

	line.Container = components.NewDefaultContainer(line)
	return line.Container

}

func (line *Line) GetPage() *components.Page {
	if line.Page != nil {
		return line.Page
	}

	line.Page = components.NewDefaultPage(line.GetContainer())
	return line.Page
}

func (line *Line) Render(file string) {
	doRender(file, line.GetPage())
}

func NewLine() *Line {
	line := &Line{}

	line.LineConfiguration = &LineConfiguration{
		BaseConfiguration: config.BaseConfiguration{}.New(),
		XAxis:             opts.XAxis{}.New(),
		YAxis:             opts.YAxis{}.New(),
		Series:            series.LineSeries{Type: types.ChartLine}.New(),
	}

	line.Container = components.NewDefaultContainer(line)
	line.Page = components.NewDefaultPage(line.Container)

	return line
}
