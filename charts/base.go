package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"regexp"

	"github.com/go-echarts/go-echarts/datatypes"
)

type GlobalOpts func(bc *BaseConfiguration)

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

// BaseOpts represents a option set needed by all chart types.
type BaseConfiguration struct {
	opts.Initialization // 图形初始化配置项
	opts.Legend         // 图例组件配置项
	opts.Tooltip        // 提示框组件配置项
	opts.Toolbox        // 工具箱组件配置项
	opts.Title          // 标题组件配置项

	AssetsOpts            // 静态资源配置项
	legends               []string
	Colors                []string // 全局颜色列表
	appendColor           []string // 追加全局颜色列表
	Routers                        // 路由列表
	DataZoomOptsList               // 区域缩放组件配置项列表
	VisualMapOptsList              // 视觉映射组件配置项列表
	RadarComponentOpts             // 雷达图组件配置项
	ParallelComponentOpts          // 平行坐标系组件配置项
	ParallelAxisOpts               // 平行坐标系中的坐标轴组件配置项
	JSFunctions                    // JS 函数列表
	opts.SingleAxis                // 单轴组件

	// todo: 使用方法代替属性
	HasXYAxis     bool // 图形是否拥有 XY 轴
	Has3DAxis     bool // 图形是否拥有 3D XYZ 轴
	HasGeo        bool // 图形是否拥有 Geo 组件
	HasRadar      bool // 图形是否拥有 Radar 组件
	HasParallel   bool // 图形是否拥有 Parallel 组件
	HasSingleAxis bool // 图形是否拥有 singleAxis 组件
}

// 设置全局颜色
func (opt *BaseOpts) setColor(options ...SeriesOptser) {
	for i := 0; i < len(options); i++ {
		switch option := options[i].(type) {
		case ColorOpts:
			opt.insertSeriesColors(option)
		}
	}
}

// 初始化全局颜色列表
func (opt *BaseOpts) initSeriesColors() {
	opt.Colors = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae",
		"#749f83", "#ca8622", "#bda29a", "#6e7074", "#546570",
	}
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
func (opt *BaseOpts) setBaseGlobalOptions(options ...GlobalOptser) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option := option.(type) {
		case InitOpts:
			opt.InitOpts = option
			if opt.InitOpts.Theme != "" {
				opt.JSAssets.Add("themes/" + opt.Theme + ".js")
			}
			opt.validateInitOpt()
		case TitleOpts:
			opt.TitleOpts = option
		case ToolboxOpts:
			opt.ToolboxOpts = option
		case TooltipOpts:
			opt.TooltipOpts = option
		case LegendOpts:
			opt.LegendOpts = option
		case ColorOpts:
			opt.insertSeriesColors(option)
		case DataZoomOpts:
			opt.DataZoomOptsList = append(opt.DataZoomOptsList, option)
		case VisualMapOpts:
			opt.VisualMapOptsList = append(opt.VisualMapOptsList, option)
		case RadarComponentOpts:
			opt.RadarComponentOpts = option
		case ParallelComponentOpts:
			opt.ParallelComponentOpts = option
		case ParallelAxisOpts:
			opt.ParallelAxisOpts = option
		case SingleAxisOpts:
			opt.SingleAxisOpts = option
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
