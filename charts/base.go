package charts

import (
	"github.com/go-echarts/go-echarts/opts"
)

// GlobalOpts sets tje Global options for charts.
type GlobalOpts func(bc *BaseConfiguration)

// BaseConfiguration represents a option set needed by all chart types.
type BaseConfiguration struct {
	opts.Initialization    // 图形初始化配置项
	opts.Legend            // 图例组件配置项
	opts.Tooltip           // 提示框组件配置项
	opts.Toolbox           // 工具箱组件配置项
	opts.Title             // 标题组件配置项
	opts.Assets            // 静态资源配置项
	opts.RadarComponent    // 雷达图组件配置项
	opts.ParallelComponent // 平行坐标系组件配置项
	opts.JSFunctions       // JS 函数列表
	opts.SingleAxis        // 单轴组件

	HasXYAxis bool // 图形是否拥有 XY 轴
	XYAxis

	Has3DAxis bool // 图形是否拥有 3D XYZ 轴
	opts.XAxis3D
	opts.YAxis3D
	opts.ZAxis3D
	opts.Grid3D

	legends     []string
	Colors      []string // 全局颜色列表
	appendColor []string // 追加全局颜色列表

	Routers          []opts.Router       // 路由列表
	DataZoomList     []opts.DataZoom     // 区域缩放组件配置项列表
	VisualMapList    []opts.VisualMap    // 视觉映射组件配置项列表
	ParallelAxisList []opts.ParallelAxis // 平行坐标系中的坐标轴组件配置项

	HasGeo        bool // 图形是否拥有 Geo 组件
	HasRadar      bool // 图形是否拥有 Radar 组件
	HasParallel   bool // 图形是否拥有 Parallel 组件
	HasSingleAxis bool // 图形是否拥有 singleAxis 组件
}

func (bc *BaseConfiguration) initBaseConfiguration() {
	bc.initSeriesColors()
	bc.InitAssets()
	bc.Initialization.Validate()
}

func (bc *BaseConfiguration) initSeriesColors() {
	bc.Colors = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae",
		"#749f83", "#ca8622", "#bda29a", "#6e7074", "#546570",
	}
}

func (bc *BaseConfiguration) insertSeriesColors(cs opts.Colors) {
	tmpCl := reverseSlice(cs)
	for i := 0; i < len(tmpCl); i++ {
		bc.Colors = append(bc.Colors, "")
		copy(bc.Colors[1:], bc.Colors[0:])
		bc.Colors[0] = tmpCl[i]
	}
}

func (bc *BaseConfiguration) setBaseGlobalOptions(opts ...GlobalOpts) {
	for _, opt := range opts {
		opt(bc)
	}
}

// WithTitleOpts
func WithTitleOpts(opt opts.Title) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Title = opt
	}
}

// WithToolboxOpts
func WithToolboxOpts(opt opts.Toolbox) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Toolbox = opt
	}
}

// WithTooltipOpts
func WithTooltipOpts(opt opts.Tooltip) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Tooltip = opt
	}
}

// WithLegendOpts
func WithLegendOpts(opt opts.Legend) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Legend = opt
	}
}

// WithInitializationOpts
func WithInitializationOpts(opt opts.Initialization) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Initialization = opt
		if bc.Initialization.Theme != "" {
			bc.JSAssets.Add("themes/" + opt.Theme + ".js")
		}
		bc.Initialization.Validate()
	}
}

// WithDataZoomOpts
func WithDataZoomOpts(opt ...opts.DataZoom) GlobalOpts {
	return func(bc *BaseConfiguration) {
		for _, o := range opt {
			bc.DataZoomList = append(bc.DataZoomList, o)
		}
	}
}

// WithVisualMapOpts
func WithVisualMapOpts(opt ...opts.VisualMap) GlobalOpts {
	return func(bc *BaseConfiguration) {
		for _, o := range opt {
			bc.VisualMapList = append(bc.VisualMapList, o)
		}
	}
}

// WithRadarComponentOpts
func WithRadarComponentOpts(opt opts.RadarComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.RadarComponent = opt
	}
}

// WithParallelComponentOpts
func WithParallelComponentOpts(opt opts.ParallelComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelComponent = opt
	}
}

// WithColorsOpts
func WithColorsOpts(opt opts.Colors) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.insertSeriesColors(opt)
	}
}

// WithRouterOpts
func WithRouterOpts(opt opts.Router) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Routers = append(bc.Routers, opt)
	}
}

// reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
