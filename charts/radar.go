package charts

import (
	"io"
)

// IndicatorOpts is the option set for a radar chart.
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

// RadarComponentOpts is the option set for a radar component.
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

// Radar represents a radar chart.
type Radar struct {
	BaseOpts
	Series
}

func (Radar) chartType() string { return ChartType.Radar }

// NewRadar creates a new radar chart.
func NewRadar(routers ...RouterOpts) *Radar {
	chart := new(Radar)
	chart.initBaseOpts(routers...)
	chart.HasRadar = true
	return chart
}

// Add adds new data sets.
func (c *Radar) Add(name string, data interface{}, options ...seriesOptser) *Radar {
	series := singleSeries{Name: name, Type: ChartType.Radar, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	c.legends = append(c.legends, name)
	return c
}

// SetGlobalOptions sets options for the Radar instance.
func (c *Radar) SetGlobalOptions(options ...globalOptser) *Radar {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Radar) validateOpts() {
	c.LegendOpts.Data = c.legends
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Radar) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
