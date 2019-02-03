package charts

import (
	"io"
)

type Pie struct {
	BaseOpts
	Series
}

func (Pie) chartType() string { return "pie" }

// Pie series options
type PieOpts struct {
	RoseType interface{}
	Center   interface{}
	Radius   interface{}
}

func (PieOpts) markSeries() {}

func (opt *PieOpts) setChartOpt(s *singleSeries) {
	s.RoseType = opt.RoseType
	s.Center = opt.Center
	s.Radius = opt.Radius
}

func NewPie(routers ...HTTPRouter) *Pie {
	chart := new(Pie)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Pie) Add(name string, data map[string]interface{}, options ...seriesOptser) *Pie {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "pie", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Pie) SetGlobalOptions(options ...globalOptser) *Pie {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Pie) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Pie) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
