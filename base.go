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

// 图形初始化配置项
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"800px"`
	// 画布高度
	Height string `default:"500px"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// JS host 地址
	JSHost string `default:"https://cdn.bootcss.com"`
}

// 为 InitOptions 设置字段默认值
func (opt *InitOpts) setDefault() {
	err := setDefaultValue(opt)
	checkError(err)
}

// 确保 ContainerID 不为空且唯一
func (opt *InitOpts) checkID() {
	if opt.ChartID == "" {
		opt.ChartID = genChartID()
	}
}

// 验证初始化参数，确保图形能够得到正确渲染
func (opt *InitOpts) verifyInitOpt() {
	opt.setDefault()
	opt.checkID()
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

// 全局颜色配置项
type ColorOpts struct {
	Color []string
}

type HttpRouter struct {
	Url  string
	Text string
}

// 所有图表都拥有的基本配置项
type BaseOpts struct {
	// 图形初始化配置项
	InitOpts
	// 图例组件配置项
	LegendOpts
	// 提示框组件配置项
	TooltipOpts
	// 标题组件配置项
	TitleOpts
	// 全局颜色列表
	ColorList []string
	// 追加全局颜色列表
	appendColor []string
	// 路由列表
	HttpRouters
}

type HttpRouters []HttpRouter

func (hr HttpRouters) Len() int {
	return len(hr)
}

// 设置全局颜色
func (opt *BaseOpts) setColor(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case ColorOpts:
			opt.appendColor = append(opt.appendColor, option.(ColorOpts).Color...)
		}
	}
}

// 初始化全局颜色列表
func (opt *BaseOpts) initSeriesColors() {
	opt.ColorList = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae", "#749f83",
		"#ca8622", "#bda29a", "#6e7074", "#546570", "#c4ccd3"}
}

//
func (opt *BaseOpts) init(routers ...HttpRouter) {
	for i := 0; i < len(routers); i++ {
		opt.HttpRouters = append(opt.HttpRouters, routers[i])
	}
	opt.initSeriesColors()
}

// 插入颜色到颜色列表首部
func (opt *BaseOpts) insertSeriesColors(s []string) {
	// 翻转颜色列表
	tmpCl := reverseSlice(s)
	// 颜色追加至首部
	for i := 0; i < len(tmpCl); i++ {
		opt.ColorList = append(opt.ColorList, "")
		copy(opt.ColorList[1:], opt.ColorList[0:])
		opt.ColorList[0] = tmpCl[i]
	}
}

// 设置 BaseOptions 全局配置项
func (opt *BaseOpts) setBaseGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOpts:
			opt.InitOpts = option.(InitOpts)
		case TitleOpts:
			opt.TitleOpts = option.(TitleOpts)
		case LegendOpts:
			opt.LegendOpts = option.(LegendOpts)
		case ColorOpts:
			opt.insertSeriesColors(option.(ColorOpts).Color)
		}
	}
}

//TODO: dataZoom
//TODO: visualMap

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
