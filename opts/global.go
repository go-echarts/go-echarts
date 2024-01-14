package opts

import (
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/types"
)

const (
	EchartsJS = "echarts.min.js"
	// CompatibleEchartsJS The 3d charts and 3rd charts not support in v5+ echarts version, back to v4 (v4.9.0)
	CompatibleEchartsJS = "echarts@4.min.js"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Initialization contains options for the canvas.
type Initialization struct {
	// HTML title
	PageTitle string `default:"Awesome go-echarts"`

	// Width of canvas
	Width string `default:"900px"`

	// Height of canvas
	Height string `default:"500px"`

	// BackgroundColor of canvas
	BackgroundColor string

	// Chart unique ID
	ChartID string

	// Assets host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`

	// Theme of chart
	Theme string `default:"white"`
}

// Validate validates the initialization configurations.
func (opt *Initialization) Validate() {
	SetDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = generateUniqueID()
	}
}

// SetDefaultValue set default values for the struct field.
// inspired from: https://github.com/mcuadros/go-defaults
func SetDefaultValue(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	walkField(elem)
}

func walkField(val reflect.Value) {
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		f := val.Field(i)
		if f.Kind() == reflect.Struct {
			walkField(f)
		}

		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(val.Field(i), defaultVal)
		}
	}
}

// setField handles String/Bool types only.
func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}

const (
	chartIDSize = 12
)

// generate the unique ID for each chart.
func generateUniqueID() string {
	var b [chartIDSize]byte
	for i := range b {
		b[i] = randByte()
	}
	return string(b[:])
}

func randByte() byte {
	c := 65 // A
	if rand.Intn(10) > 5 {
		c = 97 // a
	}
	return byte(c + rand.Intn(26))
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

// Used to customize how to slice continuous data, and some specific view style for some pieces.
type Piece struct {
	Min float32 `json:"min,omitempty"`

	Max float32 `json:"max,omitempty"`

	Lt float32 `json:"lt,omitempty"`

	Lte float32 `json:"lte,omitempty"`

	Gt float32 `json:"gt,omitempty"`

	Gte float32 `json:"gte,omitempty"`

	// Symbol color
	Color string `json:"color,omitempty"`
}

// Indicator is the option set for a radar chart.
type Indicator struct {
	// Indicator name
	Name string `json:"name,omitempty"`

	// The maximum value of indicator. It is an optional configuration, but we recommend to set it manually.
	Max float32 `json:"max,omitempty"`

	// The minimum value of indicator. It it an optional configuration, with default value of 0.
	Min float32 `json:"min,omitempty"`

	// Specify a color the the indicator.
	Color string `json:"color,omitempty"`
}

var newlineTabPat = regexp.MustCompile(`\n|\t`)
var commentPat = regexp.MustCompile(`(//.*)\n`)

type Colors []string

// Assets contains options for static assets.
type Assets struct {
	JSAssets  types.OrderedSet
	CSSAssets types.OrderedSet

	CustomizedJSAssets  types.OrderedSet
	CustomizedCSSAssets types.OrderedSet
}

// InitAssets inits the static assets storage.
func (opt *Assets) InitAssets() {
	opt.JSAssets.Init(EchartsJS)
	opt.CSSAssets.Init()

	opt.CustomizedJSAssets.Init()
	opt.CustomizedCSSAssets.Init()
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
