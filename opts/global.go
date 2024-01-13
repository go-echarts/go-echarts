package opts

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// Brush is an area-selecting component, with which user can select part of data from a chart to display in detail, or do calculations with them.
// https://echarts.apache.org/en/option.html#brush
type Brush struct {

	//XAxisIndex Assigns which of the xAxisIndex can use brush selecting.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	//Brushlink is a mapping of dataIndex. So data of every series with brushLink should be guaranteed to correspond to the other.
	Brushlink interface{} `json:"brushlink,omitempty"`

	//OutOfBrush Defines visual effects of items out of selection
	OutOfBrush *BrushOutOfBrush `json:"outOfBrush,omitempty"`
}

// BrushOutOfBrush
// https://echarts.apache.org/en/option.html#brush.outOfBrush
type BrushOutOfBrush struct {
	ColorAlpha float32 `json:"colorAlpha,omitempty"`
}

// Toolbox is the option set for a toolbox component.
// https://echarts.apache.org/en/option.html#toolbox
type Toolbox struct {
	// Whether to show toolbox component.
	Show types.Bool `json:"show,omitempty"`

	// The layout orientation of toolbox's icon.
	// Options: 'horizontal','vertical'
	Orient string `json:"orient,omitempty"`

	// Distance between toolbox component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between toolbox component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between toolbox component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between toolbox component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// The configuration item for each tool.
	// Besides the tools we provide, user-defined toolbox is also supported.
	Feature *ToolBoxFeature `json:"feature,omitempty"`
}

// ToolBoxFeature is a feature component under toolbox.
// https://echarts.apache.org/en/option.html#toolbox
type ToolBoxFeature struct {
	// Save as image tool
	SaveAsImage *ToolBoxFeatureSaveAsImage `json:"saveAsImage,omitempty"`

	// Data brush
	Brush *ToolBoxFeatureBrush `json:"brush,omitempty"`

	// Data area zooming, which only supports rectangular coordinate by now.
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`

	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`

	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`

	// User-defined tools. They have to start with "my".
	UserDefined map[string]ToolBoxFeatureUserDefined `json:"-"`
}

func (f ToolBoxFeature) MarshalJSON() ([]byte, error) {
	type ToolBoxFeature_ ToolBoxFeature

	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)

	err := enc.Encode(ToolBoxFeature_(f))
	if err != nil {
		return nil, err
	}
	// Remove trailing newline.
	buff.Truncate(buff.Len() - 1)

	if len(f.UserDefined) == 0 {
		return buff.Bytes(), nil
	}

	user := new(bytes.Buffer)
	enc = json.NewEncoder(user)
	enc.SetEscapeHTML(false)

	err = enc.Encode(f.UserDefined)
	if err != nil {
		return nil, err
	}
	// Remove trailing newline.
	user.Truncate(user.Len() - 1)

	// Remove trailing "}".
	buff.Truncate(buff.Len() - 1)
	_, _ = buff.WriteString(",")
	// Remove prefix "{".
	_, _ = user.ReadByte()
	// Copy user-defined tools over.
	_, _ = buff.ReadFrom(user)
	return buff.Bytes(), nil
}

// ToolBoxFeatureUserDefined is the option fro user-defined tools.
// https://echarts.apache.org/en/option.html#toolbox.feature
type ToolBoxFeatureUserDefined struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`

	// Icon for the tool.
	Icon string `json:"icon,omitempty"`

	// On click handler in JavaScript. Use opts.FuncOpts to embed JavaScript.
	OnClick string `json:"onclick,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
// https://echarts.apache.org/en/option.html#toolbox.feature.saveAsImage
type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default), t
	// hen 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when when chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).
	Type string `json:"type,omitempty" default:"png"`

	// Name to save the image, whose default value is title.text.
	Name string `json:"name,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`
}

// ToolBoxFeatureBrush  brush-selecting icon.
// https://echarts.apache.org/en/option.html#toolbox.feature.brush
type ToolBoxFeatureBrush struct {

	//Icons used, whose values are:
	// 'rect': Enabling selecting with rectangle area.
	// 'polygon': Enabling selecting with any shape.
	// 'lineX': Enabling horizontal selecting.
	// 'lineY': Enabling vertical selecting.
	// 'keep': Switching between single selecting and multiple selecting. The latter one can select multiple areas, while the former one cancels previous selection.
	// 'clear': Clearing all selection.
	Type []string `json:"type,omitempty"`
}

// ToolBoxFeatureDataZoom
// https://echarts.apache.org/en/option.html#toolbox.feature.dataZoom
type ToolBoxFeatureDataZoom struct {
	// Whether to show the tool.
	Show types.Bool `json:"show"`

	//Defines which yAxis should be controlled. By default, it controls all y axes.
	//If it is set to be false, then no y axis is controlled.
	//If it is set to be then it controls axis with axisIndex of 3.
	//If it is set to be [0, 3], it controls the x-axes with axisIndex of 0 and 3.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`

	// Restored and zoomed title text.
	// m["zoom"] = "area zooming"
	// m["back"] = "restore area zooming"
	Title map[string]string `json:"title"`
}

// ToolBoxFeatureDataView
// https://echarts.apache.org/en/option.html#toolbox.feature.dataView
type ToolBoxFeatureDataView struct {
	// Whether to show the tool.
	Show types.Bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`

	// There are 3 names in data view
	// you could set them like this: []string["data view", "turn off", "refresh"]
	Lang []string `json:"lang"`

	// Background color of the floating layer in data view.
	BackgroundColor string `json:"backgroundColor"`
}

// ToolBoxFeatureRestore
// https://echarts.apache.org/en/option.html#toolbox.feature.restore
type ToolBoxFeatureRestore struct {
	// Whether to show the tool.
	Show types.Bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}

// AxisLabel settings related to axis label.
// https://echarts.apache.org/en/option.html#xAxis.axisLabel
type AxisLabel struct {
	// Set this to false to prevent the axis label from appearing.
	Show types.Bool `json:"show"`

	// Interval of Axis label, which is available in category axis.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label.
	// And if it is set to be 2, it means labels are shown once after two labels, and so on.
	Interval string `json:"interval,omitempty"`

	// Set this to true so the axis labels face the inside direction.
	Inside types.Bool `json:"inside,omitempty"`

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

	ShowMinLabel types.Bool `json:"showMinLabel"`
	ShowMaxLabel types.Bool `json:"showMaxLabel"`

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

	BackgroundColor string `json:"backgroundColor,omitempty"`
}

type AxisTick struct {
	// Set this to false to prevent the axis tick from showing.
	Show types.Bool `json:"show"`

	// interval of axisTick, which is available in category axis. is set to be the same as axisLabel.interval by default.
	// It uses a strategy that labels do not overlap by default.
	// You may set it to be 0 to display all labels compulsively.
	// If it is set to be 1, it means that labels are shown once after one label. And if it is set to be 2, it means labels are shown once after two labels, and so on.
	// On the other hand, you can control by callback function, whose format is shown below:
	// (index:number, value: string) => types.Boolean
	// The first parameter is index of category, and the second parameter is the name of category. The return values decides whether to display label.
	Interval string `json:"interval,omitempty"`

	// Align axis tick with label, which is available only when boundaryGap is set to be true in category axis.
	AlignWithLabel types.Bool `json:"alignWithLabel,omitempty"`
}

// AxisLine controls settings related to axis line.
// https://echarts.apache.org/en/option.html#yAxis.axisLine
type AxisLine struct {
	// Set this to false to prevent the axis line from showing.
	Show types.Bool `json:"show"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	OnZero types.Bool `json:"onZero,omitempty"`

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
	Show types.Bool `json:"show"`

	// Split area style.
	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show types.Bool `json:"show"`

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

// SingleAxis is the option set for single axis.
// https://echarts.apache.org/en/option.html#singleAxis
type SingleAxis struct {
	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Distance between grid component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
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

// GeoComponent is the option set for geo component.
// https://echarts.apache.org/en/option.html#geo
type GeoComponent struct {
	// Map charts.
	Map string `json:"map,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Set this to true, to prevent interaction with the axis.
	Silent types.Bool `json:"silent,omitempty"`
}

// ParallelComponent is the option set for parallel component.
type ParallelComponent struct {
	// Distance between parallel component and the left side of the container.
	// Left value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%';
	// and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between parallel component and the top side of the container.
	// Top value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	// and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between parallel component and the right side of the container.
	// Right value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between parallel component and the bottom side of the container.
	// Bottom value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// ParallelAxis is the option set for a parallel axis.
type ParallelAxis struct {
	// Dimension index of coordinate axis.
	Dim int `json:"dim,omitempty"`

	// Name of axis.
	Name string `json:"name,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	// In category axis, it can also be set as the ordinal number.
	Max interface{} `json:"max,omitempty"`

	// Compulsively set segmentation interval for axis.
	Inverse types.Bool `json:"inverse,omitempty"`

	// Location of axis name. Options: "start", "middle", "center", "end"
	// default "end"
	NameLocation string `json:"nameLocation,omitempty"`

	// Type of axis.
	// Options：
	// "value"：Numerical axis, suitable for continuous data.
	// "category" Category axis, suitable for discrete category data. Category data can be auto retrieved from series.
	// "log"  Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Category data，works on (type: "category").
	Data interface{} `json:"data,omitempty"`
}

// CalendarLabel is the option set for a calendar label: DayLabel, MonthLabel, YearLabel.
type CalendarLabel struct {
	// Whether to show the label.
	Show types.Bool `json:"show"`

	// The margin between the month label and the axis line.
	Margin float64 `json:"margin,omitempty"`

	// Position of year.
	// Default: when orient is set as horizontal, position is left when orient is set as vertical, position is top.
	// Options: 'left', 'right', 'top', 'bottom'
	Position string `json:"position,omitempty"`

	// Text color.
	Color string `json:"color,omitempty"`

	// Font style.
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`

	// Font weight.
	// Options: 'normal', 'bold', 'bolder', 'lighter', 100 | 200 | 300 | 400...
	FontWeight string `json:"fontWeight,omitempty"`

	// Font family.
	FontFamily string `json:"fontFamily,omitempty"`

	// Font size.
	FontSize int `json:"fontSize,omitempty"`

	// Horizontal alignment of text, automatic by default.
	// Options: 'left', 'center', 'right'
	Align string `json:"align,omitempty"`

	// Vertical alignment of text, automatic by default.
	// Options: 'top', 'middle', 'bottom'
	VerticalAlign string `json:"verticalAlign,omitempty"`

	// Line height of text.
	LineHeight int `json:"lineHeight,omitempty"`

	// Background color of label, which is transparent by default.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Border color of label.
	BorderColor string `json:"borderColor,omitempty"`

	// Border width of label.
	BorderWidth int `json:"borderWidth,omitempty"`

	// Border radius of label.
	BorderRadius int `json:"borderRadius,omitempty"`

	// Border type of label.
	// Options: 'solid', 'dashed', 'dotted'
	BorderType string `json:"borderType,omitempty"`

	// Border dash offset of label.
	BorderDashOffset int `json:"borderDashOffset,omitempty"`

	// Padding
	Padding int `json:"padding,omitempty"`

	// Shadow blur of text block.
	ShadowBlur int `json:"shadowBlur,omitempty"`

	// Shadow color of text block.
	ShadowColor string `json:"shadowColor,omitempty"`

	// Shadow X offset of text block.
	ShadowOffsetX int `json:"shadowOffsetX,omitempty"`

	// Shadow Y offset of text block.
	ShadowOffsetY int `json:"shadowOffsetY,omitempty"`

	// Width
	Width int `json:"width,omitempty"`

	// Height
	Height int `json:"height,omitempty"`

	// Text border color.
	TextBorderColor string `json:"textBorderColor,omitempty"`

	// Text border width.
	TextBorderWidth int `json:"textBorderWidth,omitempty"`

	// Text border type
	// Options: 'solid', 'dashed', 'dotted'
	TextBorderType string `json:"textBorderType,omitempty"`

	// Text border dash offset.
	TextBorderDashOffset int `json:"textBorderDashOffset,omitempty"`

	// Text shadow color.
	TextShadowColor string `json:"textShadowColor,omitempty"`

	// Text shadow blur.
	TextShadowBlur int `json:"textShadowBlur,omitempty"`

	// Text shadow X offset.
	TextShadowOffsetX int `json:"textShadowOffsetX,omitempty"`

	// Text shadow Y offset.
	TextShadowOffsetY int `json:"textShadowOffsetY,omitempty"`

	// Overflow
	Overflow string `json:"overflow,omitempty"`

	// Ellipsis
	Ellipsis types.Bool `json:"ellipsis,omitempty"`

	// Silent
	Silent types.Bool `json:"silent,omitempty"`
}

// Calendar is the option set for a calendar component.
// This works with the calendar coordinate system, and it is a heatmap calendar.
type Calendar struct {
	ID     string `json:"id,omitempty"`
	Zlevel int    `json:"zlevel,omitempty"`
	Z      int    `json:"z,omitempty"`
	// Distance between grid component and the left side of the container.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the right side of the container.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the top side of the container.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the bottom side of the container.
	Bottom string `json:"bottom,omitempty"`

	// Width of grid component.
	Width string `json:"width,omitempty"`

	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Required, range of Calendar coordinates, support multiple formats.
	Range []string `json:"range,omitempty"`

	// The size of each rect of calendar coordinates.
	CellSize string `json:"cellSize,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Day Label
	DayLabel *CalendarLabel `json:"dayLabel,omitempty"`

	// Month Label
	MonthLabel *CalendarLabel `json:"monthLabel,omitempty"`

	// Year Label
	YearLabel *CalendarLabel `json:"yearLabel,omitempty"`

	// Whether to ignore mouse events. Default value is false, for triggering and responding to mouse events.
	Silent types.Bool `json:"silent,omitempty"`
}

var newlineTabPat = regexp.MustCompile(`\n|\t`)
var commentPat = regexp.MustCompile(`(//.*)\n`)

const funcMarker = "__f__"

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, newlineTabPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts returns a string suitable for options expecting JavaScript code.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

// FuncStripCommentsOpts returns a string suitable for options expecting JavaScript code,
// stripping '//' comments.
func FuncStripCommentsOpts(fn string) string {
	fn = commentPat.ReplaceAllString(fn, "")
	return replaceJsFuncs(fn)
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	fn = newlineTabPat.ReplaceAllString(fn, "")
	return fmt.Sprintf("%s%s%s", funcMarker, fn, funcMarker)
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

// XAxis3D contains options for X axis in the 3D coordinate.
type XAxis3D struct {
	// Whether to display the x-axis.
	Show types.Bool `json:"show,omitempty"`

	// The name of the axis.
	Name string `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// YAxis3D contains options for Y axis in the 3D coordinate.
type YAxis3D struct {
	// Whether to display the y-axis.
	Show types.Bool `json:"show,omitempty"`

	// The name of the axis.
	Name string `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// ZAxis3D contains options for Z axis in the 3D coordinate.
type ZAxis3D struct {
	// Whether to display the z-axis.
	Show types.Bool `json:"show,omitempty"`

	// The name of the axis.
	Name string `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// Grid3D contains options for the 3D coordinate.
type Grid3D struct {
	// Whether to show the coordinate.
	Show types.Bool `json:"show,omitempty"`

	// 3D Cartesian coordinates width
	// default 100
	BoxWidth float32 `json:"boxWidth,omitempty"`

	// 3D Cartesian coordinates height
	// default 100
	BoxHeight float32 `json:"boxHeight,omitempty"`

	// 3D Cartesian coordinates depth
	// default 100
	BoxDepth float32 `json:"boxDepth,omitempty"`

	// Rotate or scale fellows the mouse
	ViewControl *ViewControl `json:"viewControl,omitempty"`
}

// ViewControl contains options for view controlling.
type ViewControl struct {
	// Whether to rotate automatically.
	AutoRotate types.Bool `json:"autoRotate,omitempty"`

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
