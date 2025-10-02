package charts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/go-echarts/go-echarts/v2/datasets"
	"github.com/go-echarts/go-echarts/v2/event"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-echarts/go-echarts/v2/util"
)

var defaultConfigurationVisitor ConfigurationVisitor = BaseConfigurationVisitor{}

// GlobalOpts sets the Global options for charts.
type GlobalOpts func(bc *BaseConfiguration)

// BaseConfiguration represents an option set needed by all chart types.
type BaseConfiguration struct {
	opts.Legend       `json:"legend"`
	opts.Tooltip      `json:"tooltip"`
	opts.Toolbox      `json:"toolbox"`
	opts.Title        `json:"title"`
	opts.Polar        `json:"polar"`
	opts.AngleAxis    `json:"angleAxis"`
	opts.RadiusAxis   `json:"radiusAxis"`
	opts.Brush        `json:"brush"`
	*opts.AxisPointer `json:"axisPointer"`
	*opts.Aria        `json:"aria"`
	Calendar          []*opts.Calendar `json:"calendar"`

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
	opts.Grid

	legends []string
	// Colors is the color list of palette.
	// If no color is set in series, the colors would be adopted sequentially and circularly
	// from this list as the colors of series.
	Colors []string

	// Animation configs
	// Animation whether enable the animation, default true
	Animation          types.Bool `json:"animation,omitempty"`
	AnimationThreshold types.Int  `json:"animationThreshold,omitempty"`
	// AnimationDuration defined as types.FuncStr for more flexibilities, so are other related options
	AnimationDuration       types.FuncStr `json:"animationDuration,omitempty"`
	AnimationEasing         string        `json:"animationEasing,omitempty"`
	AnimationDelay          types.FuncStr `json:"animationDelay,omitempty"`
	AnimationDurationUpdate types.FuncStr `json:"animationDurationUpdate,omitempty"`
	AnimationEasingUpdate   string        `json:"animationEasingUpdate,omitempty"`
	AnimationDelayUpdate    types.FuncStr `json:"animationDelayUpdate,omitempty"`

	//Progressive specifies the amount of graphic elements that can be rendered within a frame (about 16ms) if "progressive rendering" enabled.
	//By default, progressive is auto-enabled when data amount is bigger than progressiveThreshold
	Progressive types.Int `json:"progressive,omitempty"`
	//ProgressiveThreshold number If current data amount is over the threshold, "progressive rendering" is enabled, default 3000
	ProgressiveThreshold types.Int `json:"progressiveThreshold,omitempty"`

	// Array of datasets, managed by AddDataset()
	DatasetList []opts.Dataset `json:"dataset,omitempty"`

	DataZoomList  []opts.DataZoom  `json:"datazoom,omitempty"`
	VisualMapList []opts.VisualMap `json:"visualmap,omitempty"`

	EventListeners       []event.Listener `json:"-"`
	configurationVisitor ConfigurationVisitor

	// ParallelAxisList represents the component list which is the coordinate axis for parallel coordinate.
	ParallelAxisList []opts.ParallelAxis

	has3DAxis     bool
	hasXYAxis     bool
	hasGeo        bool
	hasRadar      bool
	hasParallel   bool
	hasSingleAxis bool
	hasPolar      bool
	hasBrush      bool

	GridList []opts.Grid `json:"grid,omitempty"`
}

func (bc *BaseConfiguration) Accept(visitor ConfigurationVisitor) {
	bc.configurationVisitor = visitor
}

// JSON wraps all the options to a map so that it could be used in the base template.
// You should call `bar.Validate()` before call this method to ensure the series data set in place
// Get data in bytes
// bs, _ : = json.Marshal(bar.JSON())
func (bc *BaseConfiguration) JSON() map[string]interface{} {
	return bc.json()
}

// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
// You should call `<chart>.Validate()` before call this method to ensure the series data set in place
func (bc *BaseConfiguration) JSONNotEscaped() template.HTML {
	obj := bc.json()
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(obj)

	return template.HTML(buff.String())
}

func (bc *BaseConfiguration) json() map[string]interface{} {
	visitor := defaultConfigurationVisitor
	if bc.configurationVisitor != nil {
		visitor = bc.configurationVisitor
	}

	obj := map[string]interface{}{
		"title":   visitor.VisitTitleOpt(bc.Title),
		"legend":  visitor.VisitLegendOpt(bc.Legend),
		"tooltip": visitor.VisitTooltipOpt(bc.Tooltip),
		"series":  visitor.VisitSeriesOpt(bc.MultiSeries),
	}

	if bc.Animation != nil {
		obj["animation"] = bc.Animation
	}

	if bc.Progressive != nil {
		obj["progressive"] = bc.Animation

	}

	if bc.ProgressiveThreshold != nil {
		obj["progressiveThreshold"] = bc.ProgressiveThreshold
	}

	// if only one item, use it directly instead of an Array
	if len(bc.DatasetList) == 1 {
		obj["dataset"] = visitor.VisitDatasets(bc.DatasetList[0])
	} else if len(bc.DatasetList) > 1 {
		obj["dataset"] = visitor.VisitDatasets(bc.DatasetList...)
	}
	if bc.AxisPointer != nil {
		obj["axisPointer"] = visitor.VisitAxisPointer(bc.AxisPointer)
	}

	if bc.Aria != nil {
		obj["aria"] = visitor.VisitAria(bc.Aria)
	}

	if bc.hasPolar {
		obj["polar"] = visitor.VisitPolar(bc.Polar)
		obj["angleAxis"] = visitor.VisitAngleAxis(bc.AngleAxis)
		obj["radiusAxis"] = visitor.VisitRadiusAxis(bc.RadiusAxis)
	}

	if bc.hasGeo {
		obj["geo"] = visitor.VisitGeo(bc.GeoComponent)
	}

	if bc.hasRadar {
		obj["radar"] = visitor.VisitRadar(bc.RadarComponent)
	}

	if bc.hasParallel {
		obj["parallel"] = visitor.VisitParallel(bc.ParallelComponent)
		obj["parallelAxis"] = visitor.VisitParallelAxis(bc.ParallelAxisList)
	}

	if bc.hasSingleAxis {
		obj["singleAxis"] = visitor.VisitSingleAxis(bc.SingleAxis)
	}

	obj["toolbox"] = visitor.VisitToolbox(bc.Toolbox)

	if len(bc.DataZoomList) > 0 {
		obj["dataZoom"] = visitor.VisitDataZooms(bc.DataZoomList)
	}

	if len(bc.VisualMapList) > 0 {
		obj["visualMap"] = visitor.VisitVisualMaps(bc.VisualMapList)
	}

	if bc.hasXYAxis {
		obj["xAxis"] = visitor.VisitXAxis(bc.XAxisList)
		obj["yAxis"] = visitor.VisitYAxis(bc.YAxisList)
	}

	if bc.has3DAxis {
		obj["xAxis3D"] = visitor.VisitXAxis3D(bc.XAxis3D)
		obj["yAxis3D"] = visitor.VisitYAxis3D(bc.YAxis3D)
		obj["zAxis3D"] = visitor.VisitZAxis3D(bc.ZAxis3D)
		obj["grid3D"] = visitor.VisitGrid3D(bc.Grid3D)
	}

	if bc.Theme == "white" {
		obj["color"] = bc.Colors
	}

	if bc.Initialization.BackgroundColor != "" {
		obj["backgroundColor"] = bc.Initialization.BackgroundColor
	}

	if len(bc.GridList) > 0 {
		obj["grid"] = visitor.VisitGrid(bc.GridList)
	}

	if bc.hasBrush {
		obj["brush"] = visitor.VisitBrush(bc.Brush)
	}

	if bc.Calendar != nil {
		obj["calendar"] = visitor.VisitCalendar(bc.Calendar)
	}

	visitor.Visit(obj)
	return obj
}

// GetAssets returns the Assets options.
func (bc *BaseConfiguration) GetAssets() opts.Assets {
	return bc.Assets
}

// AddDataset adds a Dataset to this chart
func (bc *BaseConfiguration) AddDataset(dataset ...opts.Dataset) {
	bc.DatasetList = append(bc.DatasetList, dataset...)
}

// FillDefaultValues fill default values for chart options.
func (bc *BaseConfiguration) FillDefaultValues() {
	util.SetDefaultValue(bc)
}

func (bc *BaseConfiguration) initBaseConfiguration() {
	bc.initSeriesColors()
	bc.InitAssets()
	bc.initXYAxis()
	bc.Initialization.Validate()
	bc.FillDefaultValues()
}

func (bc *BaseConfiguration) initSeriesColors() {
	bc.Colors = []string{
		"#5470c6", "#91cc75", "#fac858", "#ee6666", "#73c0de",
		"#3ba272", "#fc8452", "#9a60b4", "#ea7ccc",
	}
}

func (bc *BaseConfiguration) setSeriesColors(colors []string) {
	bc.Colors = colors
}

func (bc *BaseConfiguration) setBaseGlobalOptions(opts ...GlobalOpts) {
	for _, opt := range opts {
		opt(bc)
	}
}

// WithAngleAxisOps sets the angle of the axis.
func WithAngleAxisOps(opt opts.AngleAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.AngleAxis = opt
	}
}

// WithRadiusAxisOps sets the radius of the axis.
func WithRadiusAxisOps(opt opts.RadiusAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.RadiusAxis = opt
	}
}

// WithBrush sets the Brush.
func WithBrush(opt opts.Brush) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.hasBrush = true
		bc.Brush = opt
	}
}

// WithPolarOps sets the polar.
func WithPolarOps(opt opts.Polar) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Polar = opt
		bc.hasPolar = true
		bc.hasXYAxis = false
	}
}

// WithTitleOpts sets the title.
func WithTitleOpts(opt opts.Title) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Title = opt
	}
}

// WithAnimation enable or disable the animation.
func WithAnimation(enable bool) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Animation = opts.Bool(enable)
	}
}

// WithProgressive allows to set amount of graphic elements rendered in a frame
func WithProgressive(opt int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Progressive = opts.Int(opt)
	}
}

// WithProgressiveThreshold Allows to set treshold for progressive rendering
func WithProgressiveThreshold(opt int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ProgressiveThreshold = opts.Int(opt)
	}
}

// WithToolboxOpts sets the toolbox.
func WithToolboxOpts(opt opts.Toolbox) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Toolbox = opt
	}
}

// WithSingleAxisOpts sets the single axis.
func WithSingleAxisOpts(opt opts.SingleAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.SingleAxis = opt
	}
}

// WithTooltipOpts sets the tooltip.
func WithTooltipOpts(opt opts.Tooltip) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Tooltip = opt
	}
}

// WithLegendOpts sets the legend.
func WithLegendOpts(opt opts.Legend) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Legend = opt
	}
}

func WithEventListeners(listeners ...event.Listener) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.EventListeners = append(bc.EventListeners, listeners...)
	}
}

// WithInitializationOpts sets the initialization.
func WithInitializationOpts(opt opts.Initialization) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Initialization = opt
		if bc.Initialization.Theme != "" &&
			bc.Initialization.Theme != "white" &&
			bc.Initialization.Theme != "dark" &&
			types.PresetTheme(opt.Theme) {
			bc.JSAssets.Add("themes/" + opt.Theme + ".js")
		}
		bc.Initialization.Validate()
	}
}

// WithDataZoomOpts sets the list of the zoom data.
func WithDataZoomOpts(opt ...opts.DataZoom) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.DataZoomList = append(bc.DataZoomList, opt...)
	}
}

// WithVisualMapOpts sets the List of the visual map.
func WithVisualMapOpts(opt ...opts.VisualMap) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.VisualMapList = append(bc.VisualMapList, opt...)
	}
}

// WithRadarComponentOpts sets the component of the radar.
func WithRadarComponentOpts(opt opts.RadarComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.RadarComponent = opt
	}
}

// WithGeoComponentOpts sets the geo component.
func WithGeoComponentOpts(opt opts.GeoComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.GeoComponent = opt
		mapFile, preset := datasets.PresetMapFileNames[opt.Map]
		if preset {
			bc.JSAssets.Add("maps/" + mapFile + ".js")
		}
	}
}

// WithParallelComponentOpts sets the parallel component.
func WithParallelComponentOpts(opt opts.ParallelComponent) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelComponent = opt
	}
}

// WithParallelAxisList sets the list of the parallel axis.
func WithParallelAxisList(opt []opts.ParallelAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ParallelAxisList = opt
	}
}

// WithColorsOpts sets the color.
func WithColorsOpts(opt opts.Colors) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.setSeriesColors(opt)
	}
}

// WithGridOpts sets the List of the grid.
func WithGridOpts(opt ...opts.Grid) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.GridList = append(bc.GridList, opt...)
	}
}

// WithAxisPointerOpts sets the axis pointer.
func WithAxisPointerOpts(opt *opts.AxisPointer) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.AxisPointer = opt
	}
}

func WithAriaOpts(opt *opts.Aria) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Aria = opt
	}
}
