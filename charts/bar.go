package charts

import (
	"github.com/go-echarts/go-echarts/v2/config"
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type BarConfiguration struct {
	*config.BaseConfiguration
	Series *series.BarSeries `json:"series,omitempty"`
	XAxis  *opts.XAxis       `json:"xAxis,omitempty"`
	YAxis  *opts.YAxis       `json:"yAxis,omitempty,reserved"`
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

	bar.Container = core.NewDefaultContainer(bar)
	return bar.Container

}

func (bar *Bar) GetPage() *core.Page {
	if bar.Page != nil {
		return bar.Page
	}

	bar.Page = core.NewDefaultPage(bar.GetContainer())
	return bar.Page
}

func NewBar() *Bar {
	bar := &Bar{}

	bar.BarConfiguration = &BarConfiguration{
		BaseConfiguration: config.BaseConfiguration{}.New(),
		XAxis:             opts.XAxis{}.New(),
		YAxis:             opts.YAxis{}.New(),
		Series:            series.BarSeries{}.New(),
	}

	bar.Container = core.NewDefaultContainer(bar)
	bar.Page = core.NewDefaultPage(bar.Container)

	return bar
}
