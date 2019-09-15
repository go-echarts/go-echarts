package charts

import (
	"github.com/go-echarts/go-echarts/datatypes"
	"regexp"
)

type globalOptser interface {
	markGlobal()
}

// InitOpts contains options for the canvas.
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"900px"`
	// 画布高度
	Height string `default:"500px"`
	// 画布背景颜色
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// 静态资源 host 地址
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
	// 图表主题
	Theme string `default:"white"`
}

func (InitOpts) markGlobal() {}

// AssetsOpts contains options for static assets.
type AssetsOpts struct {
	JSAssets, CSSAssets datatypes.OrderedSet
}

// 初始化静态资源配置项
func (opt *AssetsOpts) initAssetsOpts() {
	opt.JSAssets.Init("echarts.min.js")
	opt.CSSAssets.Init("bulma.min.css")
}

// 初始化静态资源配置项
func (opt *AssetsOpts) initAssetsOptsWithoutArg() {
	opt.JSAssets.Init()
	opt.CSSAssets.Init()
}

// 返回资源列表
func (opt *AssetsOpts) yieldAssets() ([]string, []string) {
	return opt.JSAssets.Values, opt.CSSAssets.Values
}

// 校验静态资源配置项，追加 host
func (opt *AssetsOpts) validateAssets(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
	}
	for j := 0; j < len(opt.CSSAssets.Values); j++ {
		opt.CSSAssets.Values[j] = host + opt.CSSAssets.Values[j]
	}
}

// 设置 InitOptions 字段默认值
func (opt *InitOpts) setDefault() {
	setDefaultValue(opt)
}

// 确保 ChartID 不为空且唯一
func (opt *InitOpts) validateChartID() {
	if opt.ChartID == "" {
		opt.ChartID = genChartID()
	}
}

// 验证初始化参数，确保图形能够得到正确渲染
func (opt *InitOpts) validateInitOpt() {
	opt.setDefault()
	opt.validateChartID()
}

// RouterOpts contains information for routing.
type RouterOpts struct {
	URL  string // 路由 URL
	Text string // 路由显示文字
}

// Routers is an array of RouterOpts.
type Routers []RouterOpts

// Len returns the count of RouterOpts elements in array.
func (hr Routers) Len() int {
	return len(hr)
}

// JSFunctions contains a set of JS functions.
type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, replaceJsFuncs(fn[i]))
	}
}

// ColorOpts contains options for color schemes.
type ColorOpts []string

func (ColorOpts) markGlobal() {}
func (ColorOpts) markSeries() {}

// BaseOpts represents a option set needed by all chart types.
type BaseOpts struct {
	InitOpts              // 图形初始化配置项
	LegendOpts            // 图例组件配置项
	legends               []string
	TooltipOpts                    // 提示框组件配置项
	ToolboxOpts                    // 工具箱组件配置项
	TitleOpts                      // 标题组件配置项
	AssetsOpts                     // 静态资源配置项
	Colors                []string // 全局颜色列表
	appendColor           []string // 追加全局颜色列表
	Routers                        // 路由列表
	DataZoomOptsList               // 区域缩放组件配置项列表
	VisualMapOptsList              // 视觉映射组件配置项列表
	RadarComponentOpts             // 雷达图组件配置项
	ParallelComponentOpts          // 平行坐标系组件配置项
	ParallelAxisOpts               // 平行坐标系中的坐标轴组件配置项
	JSFunctions                    // JS 函数列表
	SingleAxisOpts                 // 单轴组件

	HasXYAxis     bool // 图形是否拥有 XY 轴
	Has3DAxis     bool // 图形是否拥有 3D XYZ 轴
	HasGeo        bool // 图形是否拥有 Geo 组件
	HasRadar      bool // 图形是否拥有 Radar 组件
	HasParallel   bool // 图形是否拥有 Parallel 组件
	HasSingleAxis bool // 图形是否拥有 singleAxis 组件
}

// 设置全局颜色
func (opt *BaseOpts) setColor(options ...seriesOptser) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case ColorOpts:
			opt.insertSeriesColors(option.(ColorOpts))
		}
	}
}

// 初始化全局颜色列表
func (opt *BaseOpts) initSeriesColors() {
	opt.Colors = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae", "#749f83",
		"#ca8622", "#bda29a", "#6e7074", "#546570", "#c4ccd3"}
}

// 初始化 BaseOpts
func (opt *BaseOpts) initBaseOpts(routers ...RouterOpts) {
	for i := 0; i < len(routers); i++ {
		opt.Routers = append(opt.Routers, routers[i])
	}
	opt.initSeriesColors()
	opt.initAssetsOpts()
	opt.validateInitOpt()
}

// 插入颜色到颜色列表首部
func (opt *BaseOpts) insertSeriesColors(s []string) {
	tmpCl := reverseSlice(s) // 翻转颜色列表
	// 颜色追加至首部
	for i := 0; i < len(tmpCl); i++ {
		opt.Colors = append(opt.Colors, "")
		copy(opt.Colors[1:], opt.Colors[0:])
		opt.Colors[0] = tmpCl[i]
	}
}

// 设置 BaseOptions 全局配置项
func (opt *BaseOpts) setBaseGlobalOptions(options ...globalOptser) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOpts:
			opt.InitOpts = option.(InitOpts)
			if opt.InitOpts.Theme != "" {
				opt.JSAssets.Add("themes/" + opt.Theme + ".js")
			}
			opt.validateInitOpt()
		case TitleOpts:
			opt.TitleOpts = option.(TitleOpts)
		case ToolboxOpts:
			opt.ToolboxOpts = option.(ToolboxOpts)
		case TooltipOpts:
			opt.TooltipOpts = option.(TooltipOpts)
		case LegendOpts:
			opt.LegendOpts = option.(LegendOpts)
		case ColorOpts:
			opt.insertSeriesColors(option.(ColorOpts))
		case DataZoomOpts:
			opt.DataZoomOptsList = append(opt.DataZoomOptsList, option.(DataZoomOpts))
		case VisualMapOpts:
			opt.VisualMapOptsList = append(opt.VisualMapOptsList, option.(VisualMapOpts))
		case RadarComponentOpts:
			opt.RadarComponentOpts = option.(RadarComponentOpts)
		case ParallelComponentOpts:
			opt.ParallelComponentOpts = option.(ParallelComponentOpts)
		case ParallelAxisOpts:
			opt.ParallelAxisOpts = option.(ParallelAxisOpts)
		case SingleAxisOpts:
			opt.SingleAxisOpts = option.(SingleAxisOpts)
		}
	}
}

// reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	pat, _ := regexp.Compile(`\n|\t`)
	fn = pat.ReplaceAllString(fn, "")
	return "__x__" + fn + "__x__"
}
