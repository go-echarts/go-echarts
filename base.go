package goecharts

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

// Http 路由
type HttpRouter struct {
	// 路由 URL
	Url  string
	// 路由显示文字
	Text string
}

// Http 路由列表
type HttpRouters []HttpRouter

// Len() 用于 template 方法
func (hr HttpRouters) Len() int {
	return len(hr)
}

// 全局颜色配置项
type ColorOpts struct {
	Color []string
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
	// 区域缩放组件配置项列表
	DataZoomOptsList
	// 视觉映射组件配置项列表
	VisualMapOptsList
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

// 初始化 BaseOpts 操作
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
		case DataZoomOpts:
			opt.DataZoomOptsList = append(opt.DataZoomOptsList, option.(DataZoomOpts))
		case VisualMapOpts:
			opt.VisualMapOptsList = append(opt.VisualMapOptsList, option.(VisualMapOpts))
		}
	}
}
