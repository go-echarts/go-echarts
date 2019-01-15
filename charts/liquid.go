package charts

import (
	"io"
)

type Liquid struct {
	BaseOpts
	Series
}

// Liquid series options
type LiquidOpts struct {
	Shape           string
	IsShowOutline   bool
	IsWaveAnimation bool
}

type LiquidOutlineOpts struct {
	Show bool `json:"show"`
}

func (opt *LiquidOpts) setChartOpt(s *singleSeries) {
	s.Shape = opt.Shape
	s.LiquidOutlineOpts.Show = opt.IsShowOutline
	s.IsWaveAnimation = opt.IsWaveAnimation
}

func NewLiquid(routers ...HTTPRouter) *Liquid {
	chart := new(Liquid)
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("echarts-liquidfill.min.js")
	return chart
}

func (c *Liquid) Add(name string, data interface{}, options ...interface{}) *Liquid {
	series := singleSeries{Name: name, Type: "liquidFill", Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Liquid) SetGlobalConfig(options ...interface{}) *Liquid {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Liquid) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Liquid) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{`"outline":{"show":false},?`, `"waveAnimation":false,?`}, w...)
}
