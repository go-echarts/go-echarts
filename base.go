package geocharts

// 标题组件配置项
type TitleOptions struct {
	// 主标题
	Title string `json:"text"`
	// 主标题样式配置项
	TitleStyle TextStyle
	// 副标题
	Subtitle string `json:"subtext"`
	// 副标题样式配置项
	SubtitleStyle TextStyle
	// 主标题文本超链接
	Link string `json:"link"`
	// 指定窗口打开主标题超链接
	// 'self' 当前窗口打开
	// 'blank' 新窗口打开
	Target string `json:"target"`
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

// 图形初始化配置项
type InitOptions struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"800px"`
	// 画布高度
	Height string `default:"500px"`
	// 容器 ID，用于渲染标识
	ContainerID string `default:"main"`
	// JS host 地址
	JSHost      string `default:"https://cdn.bootcss.com/echarts/4.1.0"`
	// 标题组件配置项
	TitleOptions
}

// 为 InitOptions 设置字段默认值
func (opt *InitOptions) SetDefault() {
	err := SetDefaultValue(opt)
	checkError(err)
}

// 图例组件配置项
type LegendOptions struct {
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

// 为 LegendOptions 设置字段默认值
func (opt *LegendOptions) SetDefault() {
	err := SetDefaultValue(opt);
	checkError(err)
}

// 提示框组件配置项
type TooltipOptions struct {
	// 是否显示提示框
	Show bool `json:"show,omitempty"`
	// 触发类型。
	// 'item': 数据项图形触发，主要在散点图，饼图等无类目轴的图表中使用。
	// 'axis': 坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
	// 'none': 什么都不触发。
	Trigger string `json:"trigger,omitempty"`
}

// 为 TooltipOptions 设置字段默认值
func (opt *TooltipOptions) SetDefault() {
	err := SetDefaultValue(opt);
	checkError(err)
}

// 所有图表都拥有的基本配置项
type BaseOptions struct {
	// 图形初始化配置项
	InitOptions
	// 图例组件配置项
	LegendOptions
	// 提示框组件配置项
	TooltipOptions
}

// 为 BaseOptions 设置默认值
func (opt *BaseOptions) SetDefault() {
	opt.InitOptions.SetDefault()
	opt.LegendOptions.SetDefault()
	opt.TooltipOptions.SetDefault()
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
