package opts

import (
	"strings"

	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-echarts/go-echarts/v2/util"
)

const (
	EchartsJS = "echarts.min.js"
	// CompatibleEchartsJS The 3d charts and 3rd charts not support in v5+ echarts version, back to v4 (v4.9.0)
	CompatibleEchartsJS = "echarts@4.min.js"
)

type PageConfiguration struct {
	// HTML title
	PageTitle string `default:"Awesome go-echarts"`

	// Assets host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
}

// Initialization contains options for the canvas.
type Initialization struct {
	// Width of canvas
	Width string `default:"900px"`

	// Height of canvas
	Height string `default:"500px"`

	// BackgroundColor of canvas
	BackgroundColor string

	// Chart unique ID
	ChartID string

	// Theme of chart, preset themes in types.Theme<...>
	Theme string `default:"white"`

	// Renderer
	Renderer string `default:"canvas"`

	// Page configurations duplicate, a shortcut for single chart build with page settings
	PageTitle string `default:"Awesome go-echarts"`

	// Assets host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
}

// Validate validates the initialization configurations.
func (opt *Initialization) Validate() {
	util.SetDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = util.GenerateUniqueID()
	}
}

type Colors []string

// Assets contains options for static assets.
type Assets struct {
	JSAssets  types.OrderedSet
	CSSAssets types.OrderedSet

	CustomizedJSAssets  types.OrderedSet
	CustomizedCSSAssets types.OrderedSet
	CustomizedHeaders   types.OrderedSet
}

// InitAssets inits the static assets' storage.
func (opt *Assets) InitAssets() {
	opt.JSAssets.Init(EchartsJS)
	opt.CSSAssets.Init()

	opt.CustomizedJSAssets.Init()
	opt.CustomizedCSSAssets.Init()
	opt.CustomizedHeaders.Init()
}

// ClearPresetAssets clear both the preset JS and CSS static assets.
func (opt *Assets) ClearPresetAssets() {
	opt.ClearPresetJSAssets()
	opt.ClearPresetCSSAssets()
}

// ClearPresetJSAssets only clear all the preset JS static assets.
func (opt *Assets) ClearPresetJSAssets() {
	opt.JSAssets.Clear()
}

// ClearPresetCSSAssets only clear all the preset CSS static assets.
func (opt *Assets) ClearPresetCSSAssets() {
	opt.CSSAssets.Clear()
}

// AddCustomizedJSAssets adds the customized javascript assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedJSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedJSAssets.Add(assets[i])
	}
}

// AddCustomizedCSSAssets adds the customized css assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedCSSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedCSSAssets.Add(assets[i])
	}
}

// AddCustomizedHeaders adds the customized headers, should be valid with header tag, e.g.
// <script src="assets/go-echarts/example.js"></script>
func (opt *Assets) AddCustomizedHeaders(headers ...string) {
	for i := 0; i < len(headers); i++ {
		opt.CustomizedHeaders.Add(headers[i])
	}
}

// Validate validates the static assets configurations
func (opt *Assets) Validate(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		if !strings.HasPrefix(opt.JSAssets.Values[i], host) {
			opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
		}
	}

	for i := 0; i < len(opt.CSSAssets.Values); i++ {
		if !strings.HasPrefix(opt.CSSAssets.Values[i], host) {
			opt.CSSAssets.Values[i] = host + opt.CSSAssets.Values[i]
		}
	}
}

// SplitArea is the option set for a split area.
type SplitArea struct {
	// Set this to true to show the splitArea.
	Show types.Bool `json:"show,omitempty"`

	// Split area style.
	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show types.Bool `json:"show,omitempty"`

	// Split line style.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`

	// Align split line with label, which is available only when boundaryGap is set to be true in category axis.
	AlignWithLabel types.Bool `json:"alignWithLabel,omitempty"`
}
