package charts

import (
	"github.com/go-echarts/go-echarts/opts"
)

// GlobalOpts
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

	legends     []string
	Colors      []string // 全局颜色列表
	appendColor []string // 追加全局颜色列表

	Routers          []opts.Router       // 路由列表
	DataZoomList     []opts.DataZoom     // 区域缩放组件配置项列表
	VisualMapList    []opts.VisualMap    // 视觉映射组件配置项列表
	ParallelAxisList []opts.ParallelAxis // 平行坐标系中的坐标轴组件配置项

	// todo: 使用方法代替属性
	HasXYAxis     bool // 图形是否拥有 XY 轴
	Has3DAxis     bool // 图形是否拥有 3D XYZ 轴
	HasGeo        bool // 图形是否拥有 Geo 组件
	HasRadar      bool // 图形是否拥有 Radar 组件
	HasParallel   bool // 图形是否拥有 Parallel 组件
	HasSingleAxis bool // 图形是否拥有 singleAxis 组件
}

// 设置全局颜色
//func (bc *BaseConfiguration) setColor(options ...SeriesOptser) {
//	for i := 0; i < len(options); i++ {
//		switch option := options[i].(type) {
//		case ColorOpts:
//			opt.insertSeriesColors(option)
//		}
//	}
//}

// 初始化 BaseOpts
func (bc *BaseConfiguration) initBaseConfiguration(routers ...opts.Router) {
	bc.Routers = append(bc.Routers, routers...)
	bc.initSeriesColors()
	bc.InitAssets()
	//bc.ValidateAssets()
}

// 初始化全局颜色列表
func (bc *BaseConfiguration) initSeriesColors() {
	bc.Colors = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae",
		"#749f83", "#ca8622", "#bda29a", "#6e7074", "#546570",
	}
}

// 插入颜色到颜色列表首部
func (bc *BaseConfiguration) insertSeriesColors(cs opts.Colors) {
	tmpCl := reverseSlice(cs) // 翻转颜色列表
	// 颜色追加至首部
	for i := 0; i < len(tmpCl); i++ {
		bc.Colors = append(bc.Colors, "")
		copy(bc.Colors[1:], bc.Colors[0:])
		bc.Colors[0] = tmpCl[i]
	}
}

func WithTitleOpts(opt opts.Title) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Title = opt
	}
}

func WithToolboxOpts(opt opts.Toolbox) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Toolbox = opt
	}
}

func WithTooltipOpts(opt opts.Tooltip) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Tooltip = opt
	}
}

func WithLegendOpts(opt opts.Legend) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Legend = opt
	}
}

func WithInitializationOpts(opt opts.Initialization) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Initialization = opt
		if bc.Initialization.Theme != "" {
			bc.JSAssets.Add("themes/" + opt.Theme + ".js")
		}
		bc.validateInitOpt()
	}
}

func WithDataZoomOpts(opt ...opts.DataZoom) GlobalOpts {
	return func(bc *BaseConfiguration) {
		for _, o := range opt {
			bc.DataZoomList = append(bc.DataZoomList, o)
		}
	}
}

func WithVisualMapOpts(opt ...opts.VisualMap) GlobalOpts {
	return func(bc *BaseConfiguration) {
		for _, o := range opt {
			bc.VisualMapList = append(bc.VisualMapList, o)
		}
	}
}

func WithRadarComponentOpts(opt opts.RadarComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.RadarComponent = opt
	}
}

func WithParallelComponentOpts(opt opts.ParallelComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelComponent = opt
	}
}

func WithColorsOpts(opt opts.Colors) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.insertSeriesColors(opt)
	}
}

func WithRouterOpts(opt opts.Router) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Routers = append(bc.Routers, opt)
	}
}

// 设置 BaseOptions 全局配置项
func (bc *BaseConfiguration) setBaseGlobalOptions(opts ...GlobalOpts) {
	for _, opt := range opts {
		opt(bc)
	}

	//for i := 0; i < len(options); i++ {
	//	option := options[i]
	//	switch option := option.(type) {
	//case InitOpts:
	//	opt.InitOpts = option
	//	if opt.InitOpts.Theme != "" {
	//		opt.JSAssets.Add("themes/" + opt.Theme + ".js")
	//	}
	//	opt.validateInitOpt()
	//case TitleOpts:
	//	opt.TitleOpts = option
	//case ToolboxOpts:
	//	opt.ToolboxOpts = option
	//case TooltipOpts:
	//	opt.TooltipOpts = option
	//case LegendOpts:
	//	opt.LegendOpts = option
	//case ColorOpts:
	//	opt.insertSeriesColors(option)
	//case DataZoomOpts:
	//	opt.DataZoomOptsList = append(opt.DataZoomOptsList, option)
	//case VisualMapOpts:
	//	opt.VisualMapOptsList = append(opt.VisualMapOptsList, option)
	//case RadarComponentOpts:
	//	opt.RadarComponentOpts = option
	//case ParallelComponentOpts:
	//	opt.ParallelComponentOpts = option
	//case ParallelAxisOpts:
	//	opt.ParallelAxisOpts = option
	//case SingleAxisOpts:
	//	opt.SingleAxisOpts = option
	//}
	//}
}

// reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
