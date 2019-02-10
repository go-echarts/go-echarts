package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

// 雷达图配置项组件
type IndicatorOpts struct {
	// 指示器名称
	Name string `json:"name,omitempty"`
	// 指示器的最大值，可选，建议设置
	Max float32 `json:"max,omitempty"`
	// 指示器的最小值，可选，默认为 0
	Min float32 `json:"min,omitempty"`
	// 标签特定的颜色
	Color string `json:"color,omitempty"`
}

// 雷达图坐标系组件，只适用于雷达图
type RadarComponentOpts struct {
	// 雷达图的指示器，用来指定雷达图中的多个变量（维度）
	Indicator []IndicatorOpts `json:"indicator,omitempty"`
	// 雷达图绘制类型，支持 "polygon" 和 "circle"
	Shape string `json:"shape,omitempty"`
	// 指示器轴的分割段数。默认 5
	SplitNumber int `json:"splitNumber,omitempty"`
	// 坐标轴在 grid 区域中的分隔区域
	SplitArea SplitAreaOpts `json:"splitArea,omitempty"`
	// 坐标轴在 grid 区域中的分隔线
	SplitLine SplitLineOpts `json:"splitLine,omitempty"`
}

func (RadarComponentOpts) markGlobal() {}

type Radar struct {
	BaseOpts
	Series
}

func (Radar) chartType() string { return common.ChartType.Radar }

func NewRadar(routers ...RouterOpts) *Radar {
	chart := new(Radar)
	chart.initBaseOpts(false, routers...)
	chart.HasRadar = true
	return chart
}

func (c *Radar) Add(name string, data interface{}, options ...seriesOptser) *Radar {
	series := singleSeries{Name: name, Type: common.ChartType.Radar, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	c.legends = append(c.legends, name)
	return c
}

func (c *Radar) SetGlobalOptions(options ...globalOptser) *Radar {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Radar) validateOpts() {
	c.LegendOpts.Data = c.legends
	c.validateAssets(c.AssetsHost)
}

func (c *Radar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
