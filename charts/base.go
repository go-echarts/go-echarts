package charts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/go-echarts/go-echarts/v2/actions"
	"github.com/go-echarts/go-echarts/v2/datasets"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

// GlobalOpts sets the Global options for charts.
type GlobalOpts func(bc *BaseConfiguration)

// GlobalActions sets the Global actions for charts
type GlobalActions func(ba *BaseActions)

// BaseConfiguration represents an option set needed by all chart types.
type BaseConfiguration struct {
	opts.Legend       `json:"legend"`
	opts.Tooltip      `json:"tooltip"`
	opts.Toolbox      `json:"toolbox"`
	opts.Title        `json:"title"`
	opts.Dataset      `json:"dataset"`
	opts.Polar        `json:"polar"`
	opts.AngleAxis    `json:"angleAxis"`
	opts.RadiusAxis   `json:"radiusAxis"`
	opts.Brush        `json:"brush"`
	*opts.AxisPointer `json:"axisPointer"`

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
	Colors      []string
	appendColor []string // append customize color to the Colors(reverse order)

	DataZoomList  []opts.DataZoom  `json:"datazoom,omitempty"`
	VisualMapList []opts.VisualMap `json:"visualmap,omitempty"`

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

// BaseActions represents a dispatchAction set needed by all chart types.
type BaseActions struct {
	actions.Type  `json:"type,omitempty"`
	actions.Areas `json:"areas,omitempty"`
}

// JSON wraps all the options to a map so that it could be used in the base template
//
// Get data in bytes
// bs, _ : = json.Marshal(bar.JSON())
func (bc *BaseConfiguration) JSON() map[string]interface{} {
	return bc.json()
}

// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (bc *BaseConfiguration) JSONNotEscaped() template.HTML {
	obj := bc.json()
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(obj)

	return template.HTML(buff.String())
}

// JSONNotEscapedAction works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (ba *BaseActions) JSONNotEscapedAction() template.HTML {
	obj := ba.json()
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(obj)

	return template.HTML(buff.String())
}

func (bc *BaseConfiguration) json() map[string]interface{} {
	obj := map[string]interface{}{
		"title":   bc.Title,
		"legend":  bc.Legend,
		"tooltip": bc.Tooltip,
		"series":  bc.MultiSeries,
		"dataset": bc.Dataset,
	}
	if bc.AxisPointer != nil {
		obj["axisPointer"] = bc.AxisPointer
	}
	if bc.hasPolar {
		obj["polar"] = bc.Polar
		obj["angleAxis"] = bc.AngleAxis
		obj["radiusAxis"] = bc.RadiusAxis
	}

	if bc.hasGeo {
		obj["geo"] = bc.GeoComponent
	}

	if bc.hasRadar {
		obj["radar"] = bc.RadarComponent
	}

	if bc.hasParallel {
		obj["parallel"] = bc.ParallelComponent
		obj["parallelAxis"] = bc.ParallelAxisList
	}

	if bc.hasSingleAxis {
		obj["singleAxis"] = bc.SingleAxis
	}

	if bc.Toolbox.Show {
		obj["toolbox"] = bc.Toolbox
	}

	if len(bc.DataZoomList) > 0 {
		obj["dataZoom"] = bc.DataZoomList
	}

	if len(bc.VisualMapList) > 0 {
		obj["visualMap"] = bc.VisualMapList
	}

	if bc.hasXYAxis {
		obj["xAxis"] = bc.XAxisList
		obj["yAxis"] = bc.YAxisList
	}

	if bc.has3DAxis {
		obj["xAxis3D"] = bc.XAxis3D
		obj["yAxis3D"] = bc.YAxis3D
		obj["zAxis3D"] = bc.ZAxis3D
		obj["grid3D"] = bc.Grid3D
	}

	if bc.Theme == "white" {
		obj["color"] = bc.Colors
	}

	if bc.BackgroundColor != "" {
		obj["backgroundColor"] = bc.BackgroundColor
	}

	if len(bc.GridList) > 0 {
		obj["grid"] = bc.GridList
	}

	if bc.hasBrush {
		obj["brush"] = bc.Brush
	}

	return obj
}

// GetAssets returns the Assets options.
func (bc *BaseConfiguration) GetAssets() opts.Assets {
	return bc.Assets
}

// FillDefaultValues fill default values for chart options.
func (bc *BaseConfiguration) FillDefaultValues() {
	opts.SetDefaultValue(bc)
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

func (bc *BaseActions) setBaseGlobalActions(opts ...GlobalActions) {
	for _, opt := range opts {
		opt(bc)
	}
}

func (ba *BaseActions) json() map[string]interface{} {
	obj := map[string]interface{}{
		"type":  ba.Type,
		"areas": ba.Areas,
	}
	return obj
}

// WithAreas sets the areas of the action
func WithAreas(act actions.Areas) GlobalActions {
	return func(ba *BaseActions) {
		ba.Areas = act
	}
}

// WithType sets the type of the action
func WithType(act actions.Type) GlobalActions {
	return func(ba *BaseActions) {
		ba.Type = act
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
	}
}

// WithTitleOpts sets the title.
func WithTitleOpts(opt opts.Title) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Title = opt
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

// WithInitializationOpts sets the initialization.
func WithInitializationOpts(opt opts.Initialization) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Initialization = opt
		if bc.Initialization.Theme != "" &&
			bc.Initialization.Theme != "white" &&
			bc.Initialization.Theme != "dark" {
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
		bc.JSAssets.Add("maps/" + datasets.MapFileNames[opt.Map] + ".js")
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
		bc.insertSeriesColors(opt)
	}
}

// reverseSlice reverses the string slice.
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
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
