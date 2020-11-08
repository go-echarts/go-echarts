package charts

import (
	"github.com/go-echarts/go-echarts/v2/datasets"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

// GlobalOpts sets the Global options for charts.
type GlobalOpts func(bc *BaseConfiguration)

type jsonConfig func(p map[string]interface{})

// BaseConfiguration represents an option set needed by all chart types.
type BaseConfiguration struct {
	opts.Legend  `json:"legend"`
	opts.Tooltip `json:"tooltip"`
	opts.Toolbox `json:"toolbox"`
	opts.Title   `json:"title"`

	render.Renderer        `json:"-"`
	opts.Initialization    `json:"-"`
	opts.Assets            `json:"-"`
	opts.RadarComponent    `json:"-"`
	opts.GeoComponent      `json:"-"`
	opts.ParallelComponent `json:"-"`
	opts.JSFunctions       `json:"-"`
	opts.SingleAxis        `json:"-"`

	MultiSeries
	XYAxis

	opts.XAxis3D
	opts.YAxis3D
	opts.ZAxis3D
	opts.Grid3D

	legends []string
	// Colors is the color list of palette.
	// If no color is set in series, the colors would be adopted sequentially and circularly
	// from this list as the colors of series.
	Colors      []string
	appendColor []string // append customize color to the Colors(reverse order)

	DataZoomList  []opts.DataZoom  `json:"datazoom,omitempty"`
	VisualMapList []opts.VisualMap `json:"visualmap,omitempty"`

	// ParallelAxisList represents the component list which is the coordinate axis for parallel coordinate.
	ParallelAxisList []opts.ParallelAxis

	has3DAxis     bool
	hasXYAxis     bool
	hasRadar      bool
	hasParallel   bool
	hasSingleAxis bool

	jsonConfigs []jsonConfig
}

func (bc *BaseConfiguration) JSON() map[string]interface{} {
	obj := map[string]interface{}{
		"title":   bc.Title,
		"legend":  bc.Legend,
		"tooltip": bc.Tooltip,
		"series":  bc.MultiSeries,
	}

	bc.configJson(obj)

	return obj
}

func (bc *BaseConfiguration) GetAssets() opts.Assets {
	return bc.Assets
}

func (bc *BaseConfiguration) initBaseConfiguration() {
	bc.initSeriesColors()
	bc.InitAssets()
	bc.initXYAxis()
	bc.Initialization.Validate()
}

func (bc *BaseConfiguration) jsonConfig(k string, v interface{}) jsonConfig {
	return func(p map[string]interface{}) {
		p[k] = v
	}
}

func (bc *BaseConfiguration) addJSOMConfig(config ...jsonConfig) {
	bc.jsonConfigs = append(bc.jsonConfigs, config...)
}

func (bc *BaseConfiguration) configJson(data map[string]interface{}) {
	for _, config := range bc.jsonConfigs {
		config(data)
	}
}

func (bc *BaseConfiguration) initSeriesColors() {
	bc.Colors = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae",
		"#749f83", "#ca8622", "#bda29a", "#6e7074", "#546570",
	}
}

func (bc *BaseConfiguration) insertSeriesColors(colors []string) {
	reversed := reverseSlice(colors)
	for i := 0; i < len(reversed); i++ {
		bc.Colors = append(bc.Colors, "")
		copy(bc.Colors[1:], bc.Colors[0:])
		bc.Colors[0] = reversed[i]
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
		bc.Toolbox = opt // should remove this?
		bc.addJSOMConfig(bc.jsonConfig("toolbox", opt))
	}
}

// WithSingleAxisOpts
func WithSingleAxisOpts(opt opts.SingleAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.SingleAxis = opt
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

// WithGeoComponentOpts
func WithGeoComponentOpts(opt opts.GeoComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.GeoComponent = opt
		bc.JSAssets.Add("maps/" + datasets.MapFileNames[opt.Map] + ".js")
	}

}

// WithParallelComponentOpts
func WithParallelComponentOpts(opt opts.ParallelComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelComponent = opt
	}
}

// WithParallelAxisList
func WithParallelAxisList(opt []opts.ParallelAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelAxisList = opt
	}
}

// WithColorsOpts
func WithColorsOpts(opt opts.Colors) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.insertSeriesColors(opt)
	}
}

// reverseSlice reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
