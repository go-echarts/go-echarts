package charts

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineConfiguration struct {
	*opt.BaseConfiguration
	Series *series.LineSeries `json:"series,omitempty,reserved"`
	XAxis  *opt.XAxis         `json:"xAxis,omitempty,reserved"`
	YAxis  *opt.YAxis         `json:"yAxis,omitempty,reserved"`
}

// Line represents a line chart.
type Line struct {
	*LineConfiguration
}

func (line *Line) GetChartType() string {
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

func NewLine() *Line {
	line := &Line{}

	line.LineConfiguration = &LineConfiguration{
		BaseConfiguration: opt.NewBaseConfiguration(),
		XAxis:             &opt.XAxis{},
		YAxis:             &opt.YAxis{},
		Series:            &series.LineSeries{},
	}

	line.Container = core.NewContainer(line)

	return line
}

func (line *Line) AddSeries(series ...*series.LineSingleSeries) {
	for _, s := range series {
		s.Type = types.ChartLine
		*line.Series = append(*line.Series, s)
	}
}
