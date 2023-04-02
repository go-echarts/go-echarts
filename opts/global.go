package opts

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/types"
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

// set default values for the struct field.
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

// Title is the option set for a title component.
// https://echarts.apache.org/en/option.html#title
type Title struct {
	// The main title text, supporting \n for newlines.
	Title string `json:"text,omitempty"`

	// TextStyle of the main title.
	TitleStyle *TextStyle `json:"textStyle,omitempty"`

	// The hyper link of main title text.
	Link string `json:"link,omitempty"`

	// Subtitle text, supporting \n for newlines.
	Subtitle string `json:"subtext,omitempty"`

	// TextStyle of the sub title.
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`

	// The hyper link of sub title text.
	SubLink string `json:"sublink,omitempty"`

	// Open the hyper link of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target string `json:"target,omitempty"`

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`
}

// Legend is the option set for a legend component.
// Legend component shows symbol, color and name of different series. You can click legends to toggle displaying series in the chart.
// https://echarts.apache.org/en/option.html#legend
type Legend struct {
	// Whether to show the Legend, default true.
	// Once you set other options, need to manually set it to true
	Show bool `json:"show" default:"true"`

	// Type of legend. Optional values:
	// "plain": Simple legend. (default)
	// "scroll": Scrollable legend. It helps when too many legend items needed to be shown.
	Type string `json:"type"`

	// Distance between legend component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between legend component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between legend component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between legend component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Data array of legend. An array item is usually a name representing string.
	// set Data as []string{} if you wants to hide the legend.
	Data interface{} `json:"data,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Legend color when not selected.
	InactiveColor string `json:"inactiveColor,omitempty"`

	// State table of selected legend.
	// example:
	// var selected = map[string]bool{}
	// selected["series1"] = true
	// selected["series2"] = false
	Selected map[string]bool `json:"selected,omitempty"`

	// Selected mode of legend, which controls whether series can be toggled displaying by clicking legends.
	// It is enabled by default, and you may set it to be false to disabled it.
	// Besides, it can be set to 'single' or 'multiple', for single selection and multiple selection.
	SelectedMode string `json:"selectedMode,omitempty"`

	// Legend space around content. The unit is px.
	// Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	// Examples:
	// 1. Set padding to be 5
	//    padding: 5
	// 2. Set the top and bottom paddings to be 5, and left and right paddings to be 10
	//    padding: [5, 10]
	// 3. Set each of the four paddings separately
	//    padding: [
	//      5,  // up
	//      10, // right
	//      5,  // down
	//      10, // left
	//    ]
	Padding interface{} `json:"padding,omitempty"`

	// Image width of legend symbol.
	ItemWidth int `json:"itemWidth,omitempty"`

	// Image height of legend symbol.
	ItemHeight int `json:"itemHeight,omitempty"`

	// Legend X position, right/left/center
	X string `json:"x,omitempty"`

	// Legend Y position, right/left/center
	Y string `json:"y,omitempty"`

	// Width of legend component. Adaptive by default.
	Width string `json:"width,omitempty"`

	// Height of legend component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Legend marker and text aligning.
	// By default, it automatically calculates from component location and orientation.
	// When left value of this component is 'right' and orient is 'vertical', it would be aligned to 'right'.
	// Options: auto/left/right
	Align string `json:"align,omitempty"`

	// Legend text style.
	TextStyle *TextStyle `json:"textStyle,omitempty"`
}

// Tooltip is the option set for a tooltip component.
// https://echarts.apache.org/en/option.html#tooltip
type Tooltip struct {
	// Whether to show the tooltip component, including tooltip floating layer and axisPointer.
	Show bool `json:"show"`

	// Type of triggering.
	// Options:
	// * 'item': Triggered by data item, which is mainly used for charts that
	//    don't have a category axis like scatter charts or pie charts.
	// * 'axis': Triggered by axes, which is mainly used for charts that have category axes,
	//    like bar charts or line charts.
	// * 'none': Trigger nothing.
	Trigger string `json:"trigger,omitempty"`

	// Conditions to trigger tooltip. Options:
	// * 'mousemove': Trigger when mouse moves.
	// * 'click': Trigger when mouse clicks.
	// * 'mousemove|click': Trigger when mouse clicks and moves.
	// * 'none': Do not triggered by 'mousemove' and 'click'. Tooltip can be triggered and hidden
	//    manually by calling action.tooltip.showTip and action.tooltip.hideTip.
	//    It can also be triggered by axisPointer.handle in this case.
	TriggerOn string `json:"triggerOn,omitempty"`

	// Whether mouse is allowed to enter the floating layer of tooltip, whose default value is false.
	// If you need to interact in the tooltip like with links or buttons, it can be set as true.
	Enterable bool `json:"enterable,omitempty"`

	// The content formatter of tooltip's floating layer which supports string template and callback function.
	//
	// 1. String template
	// The template variables are {a}, {b}, {c}, {d} and {e}, which stands for series name,
	// data name and data value and ect. When trigger is set to be 'axis', there may be data from multiple series.
	// In this time, series index can be refereed as {a0}, {a1}, or {a2}.
	// {a}, {b}, {c}, {d} have different meanings for different series types:
	//
	// * Line (area) charts, bar (column) charts, K charts: {a} for series name,
	//   {b} for category name, {c} for data value, {d} for none;
	// * Scatter (bubble) charts: {a} for series name, {b} for data name, {c} for data value, {d} for none;
	// * Map: {a} for series name, {b} for area name, {c} for merging data, {d} for none;
	// * Pie charts, gauge charts, funnel charts: {a} for series name, {b} for data item name,
	//   {c} for data value, {d} for percentage.
	//
	// 2. Callback function
	// The format of callback function:
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// The first parameter params is the data that the formatter needs. Its format is shown as follows:
	// {
	//    componentType: 'series',
	//    // Series type
	//    seriesType: string,
	//    // Series index in option.series
	//    seriesIndex: number,
	//    // Series name
	//    seriesName: string,
	//    // Data name, or category name
	//    name: string,
	//    // Data index in input data array
	//    dataIndex: number,
	//    // Original data as input
	//    data: Object,
	//    // Value of data. In most series it is the same as data.
	//    // But in some series it is some part of the data (e.g., in map, radar)
	//    value: number|Array|Object,
	//    // encoding info of coordinate system
	//    // Key: coord, like ('x' 'y' 'radius' 'angle')
	//    // value: Must be an array, not null/undefined. Contain dimension indices, like:
	//    // {
	//    //     x: [2] // values on dimension index 2 are mapped to x axis.
	//    //     y: [0] // values on dimension index 0 are mapped to y axis.
	//    // }
	//    encode: Object,
	//    // dimension names list
	//    dimensionNames: Array<String>,
	//    // data dimension index, for example 0 or 1 or 2 ...
	//    // Only work in `radar` series.
	//    dimensionIndex: number,
	//    // Color of data
	//    color: string,
	//
	//    // the percentage of pie chart
	//    percent: number,
	// }
	Formatter string `json:"formatter,omitempty"`

	// Configuration item for axisPointer
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
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

//Brush is an area-selecting component, with which user can select part of data from a chart to display in detail, or do calculations with them.
//https://echarts.apache.org/en/option.html#brush
type Brush struct {

	//XAxisIndex Assigns which of the xAxisIndex can use brush selecting.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	//Brushlink is a mapping of dataIndex. So data of every series with brushLink should be guaranteed to correspond to the other.
	Brushlink interface{} `json:"brushlink,omitempty"`

	//OutOfBrush Defines visual effects of items out of selection
	OutOfBrush *BrushOutOfBrush `json:"outOfBrush,omitempty"`
}

//BrushOutOfBrush
//https://echarts.apache.org/en/option.html#brush.outOfBrush
type BrushOutOfBrush struct {
	ColorAlpha float32 `json:"colorAlpha,omitempty"`
}

// Toolbox is the option set for a toolbox component.
// https://echarts.apache.org/en/option.html#toolbox
type Toolbox struct {
	// Whether to show toolbox component.
	Show bool `json:"show"`

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
	Brush *ToolBoxFeatureBrush `json:"brush"`

	// Data area zooming, which only supports rectangular coordinate by now.
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`

	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`

	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
// https://echarts.apache.org/en/option.html#toolbox.feature.saveAsImage
type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default), t
	// hen 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when when chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).
	Type string `json:"png,omitempty"`

	// Name to save the image, whose default value is title.text.
	Name string `json:"name,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`
}

//ToolBoxFeatureBrush  brush-selecting icon.
//https://echarts.apache.org/en/option.html#toolbox.feature.brush
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
	Show bool `json:"show"`

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
	Show bool `json:"show"`

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
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
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

// XAxis is the option set for X axis.
// https://echarts.apache.org/en/option.html#xAxis
type XAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

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

	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// Minimum gap between split lines. For 'time' axis, MinInterval is in unit of milliseconds.
	MinInterval float64 `json:"minInterval,omitempty"`

	// Maximum gap between split lines. For 'time' axis, MaxInterval is in unit of milliseconds.
	MaxInterval float64 `json:"maxInterval,omitempty"`

	// The index of grid which the x axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of X axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Settings related to axis tick.
	AxisTick *AxisTick `json:"axisTick,omitempty"`
}

// YAxis is the option set for Y axis.
// https://echarts.apache.org/en/option.html#yAxis
type YAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

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

	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// The index of grid which the Y axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of Y axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of Y axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Settings related to axis line.
	AxisLine *AxisLine `json:"axisLine,omitempty"`
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

// VisualMap is a type of component for visual encoding, which maps the data to visual channels.
// https://echarts.apache.org/en/option.html#visualMap
type VisualMap struct {
	// Mapping type.
	// Options: "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`

	// Whether show handles, which can be dragged to adjust "selected range".
	Calculable bool `json:"calculable"`

	// Specify the min dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Min float32 `json:"min,omitempty"`

	// Specify the max dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Max float32 `json:"max,omitempty"`

	// Specify selected range, that is, the dataValue corresponding to the two handles.
	Range []float32 `json:"range,omitempty"`

	// The label text on both ends, such as ['High', 'Low'].
	Text []string `json:"text,omitempty"`

	// Specify which dimension should be used to fetch dataValue from series.data, and then map them to visual channel.
	Dimension string `json:"dimension,omitempty"`

	// Define visual channels that will mapped from dataValues that are in selected range.
	InRange *VisualMapInRange `json:"inRange,omitempty"`

	// Used to customize how to slice continuous data, and some specific view style for some pieces.
	Pieces []Piece `json:"pieces,omitempty"`

	// Whether to show visualMap-piecewise component. If set as false,
	// visualMap-piecewise component will not show,
	// but it can still perform visual mapping from dataValue to visual channel in chart.
	Show bool `json:"show"`

	// Distance between visualMap component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between visualMap component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between visualMap component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between visualMap component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// VisualMapInRange is a visual map instance in a range.
type VisualMapInRange struct {
	// Color
	Color []string `json:"color,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// DataZoom is the option set for a zoom component.
// dataZoom component is used for zooming a specific area, which enables user to
// investigate data in detail, or get an overview of the data, or get rid of outlier points.
// https://echarts.apache.org/en/option.html#dataZoom
type DataZoom struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type string `json:"type" default:"inside"`

	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	Start float32 `json:"start,omitempty"`

	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`

	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	Throttle float32 `json:"throttle,omitempty"`

	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
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

// RadarComponent is the option set for a radar component.
// https://echarts.apache.org/en/option.html#radar
type RadarComponent struct {
	// Indicator of radar chart, which is used to assign multiple variables(dimensions) in radar chart.
	Indicator []*Indicator `json:"indicator,omitempty"`

	// Radar render type, in which 'polygon' and 'circle' are supported.
	Shape string `json:"shape,omitempty"`

	// Segments of indicator axis.
	// default 5
	SplitNumber int `json:"splitNumber,omitempty"`

	// Center position of , the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width and height.
	Center interface{} `json:"center,omitempty"`

	// Split area of axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`
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
	Silent bool `json:"silent,omitempty"`
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
	Inverse bool `json:"inverse,omitempty"`

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

// Polar Bar options
type Polar struct {
	ID      string    `json:"id,omitempty"`
	Zlevel  int       `json:"zlevel,omitempty"`
	Z       int       `json:"z,omitempty"`
	Center  [2]string `json:"center,omitempty"`
	Radius  [2]string `json:"radius,omitempty"`
	Tooltip Tooltip   `json:"tooltip,omitempty"`
}

type PolarAxisBase struct {
	ID           string  `json:"id,omitempty"`
	PolarIndex   int     `json:"polarIndex,omitempty"`
	StartAngle   float64 `json:"startAngle,omitempty"`
	Type         string  `json:"type,omitempty"`
	BoundaryGap  bool    `json:"boundaryGap,omitempty"`
	Min          float64 `json:"min,omitempty"`
	Max          float64 `json:"max,omitempty"`
	Scale        bool    `json:"scale,omitempty"`
	SplitNumber  int     `json:"splitNumber,omitempty"`
	MinInterval  float64 `json:"minInterval,omitempty"`
	MaxInterval  float64 `json:"maxInterval,omitempty"`
	Interval     float64 `json:"interval,omitempty"`
	LogBase      float64 `json:"logBase,omitempty"`
	Silent       bool    `json:"silent,omitempty"`
	TriggerEvent bool    `json:"triggerEvent,omitempty"`
}

type AngleAxis struct {
	PolarAxisBase
	Clockwise bool `json:"clockwise,omitempty"`
}

type RadiusAxis struct {
	PolarAxisBase
	Name          string    `json:"name,omitempty"`
	NameLocation  string    `json:"nameLocation,omitempty"`
	NameTextStyle TextStyle `json:"nameTextStyle,omitempty"`
	NameGap       float64   `json:"nameGap,omitempty"`
	NameRadius    float64   `json:"nameRotate,omitempty"`
	Inverse       bool      `json:"inverse,omitempty"`
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

// XAxis3D contains options for X axis in the 3D coordinate.
type XAxis3D struct {
	// Whether to display the x-axis.
	Show bool `json:"show,omitempty"`

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
	Show bool `json:"show,omitempty"`

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
	Show bool `json:"show,omitempty"`

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
	Show bool `json:"show,omitempty"`

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
	AutoRotate bool `json:"autoRotate,omitempty"`

	// Rotate Speed, (angle/s).
	// default 10
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}

// Grid Drawing grid in rectangular coordinate.
// https://echarts.apache.org/en/option.html#grid
type Grid struct {
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

	ContainLabel bool `json:"containLabel,omitempty"`

	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`
}

//Dataset brings convenience in data management separated with styles and enables data reuse by different series.
//More importantly, it enables data encoding from data to visual, which brings convenience in some scenarios.
//https://echarts.apache.org/en/option.html#dataset.id
type Dataset struct {
	//source
	Source interface{} `json:"source"`
}
