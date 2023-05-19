package charts

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type BarConfiguration struct {
	*opt.BaseConfiguration
	Series *series.BarSeries `json:"series,omitempty"`
	XAxis  *opt.XAxis        `json:"xAxis,omitempty,reserved"`
	YAxis  *opt.YAxis        `json:"yAxis,omitempty,reserved"`
}

// Bar represents a bar chart.
type Bar struct {
	*BarConfiguration
}

func (bar *Bar) GetChartName() string {
	return types.ChartBar
}

func (bar *Bar) GetChart() interface{} {
	return bar
}

func (bar *Bar) GetContainer() *core.Container {
	if bar.Container != nil {
		return bar.Container
	}

	bar.Container = core.NewContainer(bar)
	return bar.Container

}

func (bar *Bar) GetPage() *core.Page {
	if bar.Page != nil {
		return bar.Page
	}

	bar.Page = core.NewPage(bar.GetContainer())
	return bar.Page
}

func NewBar() *Bar {
	bar := &Bar{}

	bar.BarConfiguration = &BarConfiguration{
		BaseConfiguration: opt.NewBaseConfiguration(),
		XAxis:             &opt.XAxis{},
		YAxis:             &opt.YAxis{},
		Series:            &series.BarSeries{},
	}

	bar.Container = core.NewContainer(bar)
	bar.Page = core.NewPage(bar.Container)

	return bar
}

func (bar *Bar) AddSeries(series ...*series.BarSingleSeries) {
	for _, s := range series {
		s.Type = types.ChartBar
		*bar.Series = append(*bar.Series, s)
	}
}
