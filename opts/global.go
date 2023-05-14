package opts

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v3/types"
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

// AxisPointer is the option set for an axisPointer component
// https://echarts.apache.org/en/option.html#axisPointer
type AxisPointer struct {

	// Indicator type.
	// Options:
	//   - 'line' line indicator.
	//   - 'shadow' shadow crosshair indicator.
	//   - 'none' no indicator displayed.
	//   - 'cross' crosshair indicator, which is actually the shortcut of enable two axisPointers of two orthometric axes.
	Type string `json:"type,omitempty"`

	// 	Whether snap to point automatically. The default value is auto determined.
	// This feature usually makes sense in value axis and time axis, where tiny points can be seeked automatically.
	Snap bool `json:"snap,omitempty"`

	Link []AxisPointerLink `json:"link,omitempty"`

	Axis string `json:"axis,omitempty"`
}

type AxisPointerLink struct {
	XAxisIndex []int  `json:"xAxisIndex,omitempty"`
	YAxisIndex []int  `json:"yAxisIndex,omitempty"`
	XAxisName  string `json:"xAxisName,omitempty"`
	YAxisName  string `json:"yAxisName,omitempty"`
}

// AxisLabel settings related to axis label.
// https://echarts.apache.org/en/option.html#xAxis.axisLabel
type AxisLabel struct {
	// Set this to false to prevent the axis label from appearing.
	Show bool `json:"show"`

	// Interval of Axis label, which is available in category axis.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label.
	// And if it is set to be 2, it means labels are shown once after two labels, and so on.
	Interval string `json:"interval,omitempty"`

	// Set this to true so the axis labels face the inside direction.
	Inside bool `json:"inside,omitempty"`

	// Rotation degree of axis label, which is especially useful when there is no enough space for category axis.
	// Rotation degree is from -90 to 90.
	Rotate float64 `json:"rotate,omitempty"`

	// The margin between the axis label and the axis line.
	Margin float64 `json:"margin,omitempty"`

	// Formatter of axis label, which supports string template and callback function.
	//
	// Example:
	//
	// Use string template; template variable is the default label of axis {value}
	// formatter: '{value} kg'
	//
	// Use callback function; function parameters are axis index
	//
	//
	//  formatter: function (value, index) {
	//    // Formatted to be month/day; display year only in the first label
	//    var date = new Date(value);
	//    var texts = [(date.getMonth() + 1), date.getDate()];
	//    if (idx === 0) {
	//        texts.unshift(date.getYear());
	//    }
	//    return texts.join('/');
	// }
	Formatter string `json:"formatter,omitempty"`

	ShowMinLabel bool `json:"showMinLabel"`
	ShowMaxLabel bool `json:"showMaxLabel"`

	// Color of axis label is set to be axisLine.lineStyle.color by default. Callback function is supported,
	// in the following format:
	//
	// (val: string) => Color
	// Parameter is the text of label, and return value is the color. See the following example:
	//
	// textStyle: {
	//    color: function (value, index) {
	//        return value >= 0 ? 'green' : 'red';
	//    }
	// }
	Color string `json:"color,omitempty"`

	// axis label font style
	FontStyle string `json:"fontStyle,omitempty"`
	// axis label font weight
	FontWeight string `json:"fontWeight,omitempty"`
	// axis label font family
	FontFamily string `json:"fontFamily,omitempty"`
	// axis label font size
	FontSize string `json:"fontSize,omitempty"`
	// Horizontal alignment of axis label
	Align string `json:"align,omitempty"`
	// Vertical alignment of axis label
	VerticalAlign string `json:"verticalAlign,omitempty"`
	// Line height of the axis label
	LineHeight string `json:"lineHeight,omitempty"`
}

type AxisTick struct {
	// Set this to false to prevent the axis tick from showing.
	Show bool `json:"show"`

	// interval of axisTick, which is available in category axis. is set to be the same as axisLabel.interval by default.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label. And if it is set to be 2, it means labels are shown once after two labels, and so on.
	// On the other hand, you can control by callback function, whose format is shown below:
	// (index:number, value: string) => boolean
	// The first parameter is index of category, and the second parameter is the name of category. The return values decides whether to display label.
	Interval string `json:"interval,omitempty"`

	// Align axis tick with label, which is available only when boundaryGap is set to be true in category axis.
	AlignWithLabel bool `json:"alignWithLabel,omitempty"`
}

// AxisLine controls settings related to axis line.
// https://echarts.apache.org/en/option.html#yAxis.axisLine
type AxisLine struct {
	// Set this to false to prevent the axis line from showing.
	Show bool `json:"show"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	OnZero bool `json:"onZero,omitempty"`

	// When multiple axes exists, this option can be used to specify which axis can be "onZero" to.
	OnZeroAxisIndex int `json:"onZeroAxisIndex,omitempty"`

	// Symbol of the two ends of the axis. It could be a string, representing the same symbol for two ends; or an array
	// with two string elements, representing the two ends separately. It's set to be 'none' by default, meaning no
	//arrow for either end. If it is set to be 'arrow', there shall be two arrows. If there should only one arrow
	//at the end, it should set to be ['none', 'arrow'].
	Symbol string `json:"symbol,omitempty"`

	// Size of the arrows at two ends. The first is the width perpendicular to the axis, the next is the width parallel to the axis.
	SymbolSize []float64 `json:"symbolSize,omitempty"`

	// Arrow offset of axis. If is array, the first number is the offset of the arrow at the beginning, and the second
	// number is the offset of the arrow at the end. If is number, it means the arrows have the same offset.
	SymbolOffset []float64 `json:"symbolOffset,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// TextStyle is the option set for a text style component.
type TextStyle struct {
	// Font color
	Color string `json:"color,omitempty"`

	// Font style
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`

	// Font size
	FontSize int `json:"fontSize,omitempty"`

	// Font family the main title font family.
	// Options: "sans-serif", 'serif' , 'monospace', 'Arial', 'Courier New', 'Microsoft YaHei', ...
	FontFamily string `json:"fontFamily,omitempty"`

	// Padding title space around content.
	// The unit is px. Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	Padding interface{} `json:"padding,omitempty"`

	// compatible for WordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}

// SplitArea is the option set for a split area.
type SplitArea struct {
	// Set this to true to show the splitArea.
	Show bool `json:"show"`

	// Split area style.
	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show bool `json:"show"`

	// Split line style.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`

	// Align split line with label, which is available only when boundaryGap is set to be true in category axis.
	AlignWithLabel bool `json:"alignWithLabel,omitempty"`
}

// Polar Bar options
type Polar struct {
	ID      string    `json:"id,omitempty"`
	Zlevel  int       `json:"zlevel,omitempty"`
	Z       int       `json:"z,omitempty"`
	Center  [2]string `json:"center,omitempty"`
	Radius  [2]string `json:"radius,omitempty"`
	Tooltip Tooltip   `json:"tooltip,omitempty"`
}

var funcPat = regexp.MustCompile(`\n|\t`)

const funcMarker = "__f__"

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, funcPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts is the option set for handling function type.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	return fmt.Sprintf("%s%s%s", funcMarker, funcPat.ReplaceAllString(fn, ""), funcMarker)
}

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
	opt.JSAssets.Init("echarts.min.js")
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

// ViewControl contains options for view controlling.
type ViewControl struct {
	// Whether to rotate automatically.
	AutoRotate bool `json:"autoRotate,omitempty"`

	// Rotate Speed, (angle/s).
	// default 10
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}

// Dataset brings convenience in data management separated with styles and enables data reuse by different series.
// More importantly, it enables data encoding from data to visual, which brings convenience in some scenarios.
// https://echarts.apache.org/en/option.html#dataset.id
type Dataset struct {
	//source
	Source interface{} `json:"source"`
}
