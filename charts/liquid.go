package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Liquid struct {
	BaseOpts
	Series
}

func (Liquid) chartType() string { return common.ChartType.Liquid }

// Liquid series options
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

type LiquidOutlineOpts struct {
	Show bool `json:"show"`
}

func (opt *LiquidOpts) setChartOpt(s *singleSeries) {
	s.Shape = opt.Shape
	s.LiquidOutlineOpts.Show = opt.IsShowOutline
	s.IsWaveAnimation = opt.IsWaveAnimation
}

func NewLiquid(routers ...RouterOpts) *Liquid {
	chart := new(Liquid)
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("echarts-liquidfill.min.js")
	return chart
}

func (c *Liquid) Add(name string, data interface{}, options ...seriesOptser) *Liquid {
	series := singleSeries{Name: name, Type: common.ChartType.Liquid, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

func (c *Liquid) SetGlobalOptions(options ...globalOptser) *Liquid {
	c.BaseOpts.setBaseGlobalOptions(options...)
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
