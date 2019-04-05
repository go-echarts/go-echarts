package charts

import (
	"io"
)

// Liquid represents a liquid chart.
type Liquid struct {
	BaseOpts
	Series
}

func (Liquid) chartType() string { return ChartType.Liquid }

// LiquidOpts is the option set for a liquid chart.
type LiquidOpts struct {
	// 水球图形状，可选
	// "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Shape string
	// 是否显示水球轮廓
	IsShowOutline bool
	// 是否停止动画
	IsWaveAnimation bool
}

func (LiquidOpts) markSeries() {}

// LiquidOutlineOpts is the options set for a liquid outline.
type LiquidOutlineOpts struct {
	Show bool `json:"show"`
}

func (opt *LiquidOpts) setChartOpt(s *singleSeries) {
	s.Shape = opt.Shape
	s.LiquidOutlineOpts.Show = opt.IsShowOutline
	s.IsWaveAnimation = opt.IsWaveAnimation
}

// NewLiquid creates a new liquid chart.
func NewLiquid(routers ...RouterOpts) *Liquid {
	chart := new(Liquid)
	chart.initBaseOpts(routers...)
	chart.JSAssets.Add("echarts-liquidfill.min.js")
	return chart
}

// Add adds new data sets.
func (c *Liquid) Add(name string, data interface{}, options ...seriesOptser) *Liquid {
	series := singleSeries{Name: name, Type: ChartType.Liquid, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

// SetGlobalOptions sets options for the Liquid instance.
func (c *Liquid) SetGlobalOptions(options ...globalOptser) *Liquid {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Liquid) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Liquid) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{`"outline":{"show":false},?`, `"waveAnimation":false,?`}, w...)
}
