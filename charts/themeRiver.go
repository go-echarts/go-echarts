package charts

import (
	"io"
)

// ThemeRiver represents a theme river chart.
type ThemeRiver struct {
	BaseOpts
	Series
}

// SingleAxisOpts is the option set for single axis.
type SingleAxisOpts struct {
	// 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度
	Min interface{} `json:"min,omitempty"`
	// 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度
	Max interface{} `json:"max,omitempty"`
	// 坐标轴类型
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "time" 时间轴，适用于连续的时序数据，与数值轴相比时间轴带有时间的格式化，
	// 在刻度计算上也有所不同，例如会根据跨度的范围来决定使用月，星期，日还是小时范围的刻度。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// single 组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Left string `json:"left,omitempty"`
	// single 组件离容器右侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Right string `json:"right,omitempty"`
	// single 组件离容器顶侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// single 组件离容器底侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Bottom string `json:"bottom,omitempty"`
}

func (SingleAxisOpts) markGlobal() {}

func (ThemeRiver) chartType() string { return ChartType.ThemeRiver }

// NewThemeRiver creates a new theme river chart.
func NewThemeRiver(routers ...RouterOpts) *ThemeRiver {
	chart := new(ThemeRiver)
	chart.initBaseOpts(routers...)
	chart.HasSingleAxis = true
	return chart
}

// Add adds new data sets.
func (c *ThemeRiver) Add(name string, data interface{}, options ...seriesOptser) *ThemeRiver {
	series := singleSeries{Name: name, Type: ChartType.ThemeRiver, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

// SetGlobalOptions sets options for the ThemeRiver instance.
func (c *ThemeRiver) SetGlobalOptions(options ...globalOptser) *ThemeRiver {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *ThemeRiver) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *ThemeRiver) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
