package charts

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineConfiguration struct {
	*opt.BaseConfiguration
	Series *series.LineSeries0 `json:"series,omitempty,reserved"`
	XAxis  *opt.XAxis          `json:"xAxis,omitempty,reserved"`
	YAxis  *opt.YAxis          `json:"yAxis,omitempty,reserved"`
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

func (line *Line) GetContainer() *core.Container {
	if line.Container != nil {
		return line.Container
	}

	line.Container = core.NewContainer(line)
	return line.Container

}

func (line *Line) GetPage() *core.Page {
	if line.Page != nil {
		return line.Page
	}

	line.Page = core.NewPage(line.GetContainer())
	return line.Page
}

func NewLine() *Line {
	line := &Line{}

	line.LineConfiguration = &LineConfiguration{
		BaseConfiguration: opt.BaseConfiguration{}.New(),
		XAxis:             opt.XAxis{}.New(),
		YAxis:             opt.YAxis{}.New(),
		Series:            &series.LineSeries0{},
	}

	line.Container = core.NewContainer(line)
	line.Page = core.NewPage(line.Container)

	return line
}

func (line *Line) AddSeries(series ...*series.LineSeries) {
	for _, s := range series {
		c := append(*line.Series, s)
		line.Series = &c
	}
}
