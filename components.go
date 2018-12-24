package goecharts

// 标题组件配置项
type TitleOpts struct {
	// 主标题
	Title string `json:"text,omitempty"`
	// 主标题样式配置项
	TitleStyle *TextStyle `json:"textStyle,omitempty"`
	// 副标题
	Subtitle string `json:"subtext,omitempty"`
	// 副标题样式配置项
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`
	// 主标题文本超链接
	Link string `json:"link,omitempty"`
	// 指定窗口打开主标题超链接
	// 'self' 当前窗口打开
	// 'blank' 新窗口打开
	Target string `json:"target,omitempty"`
	// grid 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比，也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
}

// 图例组件配置项
type LegendOpts struct {
	// 是否显示图例
	Show bool `json:"show,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比，也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
}

// 提示框组件配置项
type TooltipOpts struct {
	// 是否显示提示框
	Show bool `json:"show,omitempty"`
	// 触发类型。
	// 'item': 数据项图形触发，主要在散点图，饼图等无类目轴的图表中使用。
	// 'axis': 坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
	// 'none': 什么都不触发。
	Trigger string `json:"trigger,omitempty"`
}

// 字体样式配置项
type TextStyle struct {
	// 文字字体颜色
	Color string `json:"color,omitempty"`
	// 文字字体的风格
	// 可选  'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`
	// 字体大小
	FontSize int `json:"fontSize,omitempty"`
}

// 区域缩放组件配置项
type DataZoomOpts struct {
	// 缩放类型，可选 "inside", "slider"
	Type string `json:"type" default:"inside"`
	// 数据窗口范围的起始百分比。范围是：0 ~ 100。表示 0% ~ 100%。
	// 默认: 0
	Start float32 `json:"start,omitempty"`
	// 数据窗口范围的结束百分比。范围是：0 ~ 100。
	// 默认: 100
	End float32 `json:"end,omitempty"`
	// 触发视图刷新的频率。单位为毫秒（ms）。
	// 默认: 100
	Throttle float32 `json:"throttle,omitempty"`
	// 设置 dataZoom 组件控制的 X 轴
	// 不指定时，当 dataZoom-inside.orient 为 'horizontal'时，默认控制和 dataZoom 平行的第一个 xAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
	// 设置 dataZoom 组件控制的 Y 轴
	// 不指定时，当 dataZoom-slider.orient 为 'vertical'时，默认控制和 dataZoom 平行的第一个 yAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}

type DataZoomOptsList []DataZoomOpts

// Len() 用于 template 方法
func (dz DataZoomOptsList) Len() int {
	return len(dz)
}

// 视觉映射组件配置项
// 用于进行『视觉编码』，也就是将数据映射到视觉元素（视觉通道）
type VisualMapOpts struct {
	// 映射类型，可选 "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`
	// 是否显示拖拽用的手柄（手柄能拖拽调整选中范围）
	Calculable bool `json:"calculable"`
	// VisualMap 组件的允许的最小值
	Mix float32 `json:"mix,omitempty"`
	// VisualMap 组件的允许的最大值
	Max float32 `json:"max,omitempty"`
	// 指定手柄对应数值的位置。range 应在 min max 范围内
	Range []float32 `json:"range,omitempty"`
	// 两端的文本，如 ['High', 'Low']
	Text      []string `json:"text,omitempty"`
	*VMInRange `json:"inRange,omitempty"`
}

type VisualMapOptsList []VisualMapOpts

// Len() 用于 template 方法
func (vm VisualMapOptsList) Len() int {
	return len(vm)
}

type VMInRange struct {
	Color      []string `json:"color,omitempty"`
	Symbol     string   `json:"symbol,omitempty"`
	SymbolSize float32  `json:"symbolSize,omitempty"`
}

type VMOutRange struct {
}

// X 轴配置项组件
type XAxisOpts struct {
	// X 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 X 轴
	Show bool `json:"show,omitempty"`
	// X 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// Y 轴配置项组件
type YAxisOpts struct {
	// Y 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 Y 轴
	Show bool `json:"show,omitempty"`
	// Y 轴数据项
	Data interface{} `json:"data,omitempty"`
}
