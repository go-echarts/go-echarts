package opts

import "github.com/go-echarts/go-echarts/v2/types"

// XAxis is the option set for X axis.
// https://echarts.apache.org/en/option.html#xAxis
type XAxis struct {
	// Set this to false to prevent the axis from showing.
	Show types.Bool `json:"show,omitempty"`

	// AlignTicks turned on to automatically align ticks when multiple numeric y axes.
	// Only available for axes of type 'value' and 'log'.
	AlignTicks types.Bool `json:"alignTicks,omitempty"`

	// Position The position of x-axis.
	// options:
	// 'top'
	// 'bottom'
	Position string `json:"position,omitempty"`

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

	// Name of axis.
	Name string `json:"name,omitempty"`

	// Location of axis name.
	//
	// Options:
	// 'start'
	// 'middle' or 'center'
	// 'end'
	NameLocation string `json:"nameLocation,omitempty"`

	// Gap between axis name and axis line.
	NameGap int `json:"nameGap,omitempty"`

	// Inverse Set this to true to invert the axis.
	// Default false
	Inverse types.Bool `json:"inverse,omitempty"`

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
	Scale types.Bool `json:"scale,omitempty"`

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

	// TriggerEvent Set this to true to enable triggering events, default false
	// Parameters of the event include:
	//{
	//    // Component type: xAxis, yAxis, radiusAxis, angleAxis
	//    // Each of which has an attribute for index, e.g., xAxisIndex for xAxis
	//    componentType: string,
	//    // Value on axis before being formatted.
	//    // Click on value label to trigger event.
	//    value: '',
	//    // Name of axis.
	//    // Click on laben name to trigger event.
	//    name: ''
	//}
	TriggerEvent types.Bool `json:"triggerEvent,omitempty"`

	// The index of grid which the x axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of X axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Settings related to axis line.
	AxisLine *AxisLine `json:"axisLine,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Settings related to axis tick.
	AxisTick *AxisTick `json:"axisTick,omitempty"`

	// Settings related to axis pointer.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}

// AxisLabel settings related to axis label .
// https://echarts.apache.org/en/option.html#xAxis.axisLabel
// https://echarts.apache.org/en/option.html#yAxis.axisLabel
type AxisLabel struct {
	// Set this to false to prevent the axis label from appearing.
	Show types.Bool `json:"show,omitempty"`

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
	Formatter types.FuncStr `json:"formatter,omitempty"`

	ShowMinLabel types.Bool `json:"showMinLabel"`
	ShowMaxLabel types.Bool `json:"showMaxLabel"`

	// Alignment of the label of the min tick. If set to be null, it's the same with other labels .
	//
	// Options are:
	//
	// 'left'
	// 'center'
	// 'right'
	// null (default)
	AlignMinLabel string `json:"alignMinLabel,omitempty"`

	// Alignment of the label of the max tick. If set to be null, it's the same with other labels .
	//
	// Options are:
	//
	// 'left'
	// 'center'
	// 'right'
	// null (default)
	AlignMaxLabel string `json:"alignMaxLabel,omitempty"`

	// Whether to hide overlapped labels.
	HideOverlap types.Bool `json:"hideOverlap,omitempty"`

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
	FontSize int `json:"fontSize,omitempty"`
	// Horizontal alignment of axis label
	Align string `json:"align,omitempty"`
	// Vertical alignment of axis label
	VerticalAlign string `json:"verticalAlign,omitempty"`
	// Line height of the axis label
	LineHeight string `json:"lineHeight,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Width of text block.
	Width int `json:"width,omitempty"`

	// Height of text block.
	Height int `json:"height,omitempty"`

	// Determine how to display the text when it's overflow. Available when width is set.
	//
	// 'truncate' Truncate the text and trailing with ellipsis.
	// 'break' Break by word
	// 'breakAll' Break by character.
	Overflow string `json:"overflow,omitempty"`

	// Ellipsis to be displayed when overflow is set to truncate.
	//
	// 'truncate' Truncate the overflow lines.
	Ellipsis string `json:"ellipsis,omitempty"`
}

type AxisTick struct {
	// Set this to false to prevent the axis tick from showing.
	Show types.Bool `json:"show,omitempty"`

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

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// AxisLine controls settings related to axis line.
// https://echarts.apache.org/en/option.html#xAxis.axisLine
// https://echarts.apache.org/en/option.html#yAxis.axisLine
type AxisLine struct {
	// Set this to false to prevent the axis line from showing.
	Show types.Bool `json:"show,omitempty"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	OnZero types.Bool `json:"onZero,omitempty"`

	// When multiple axes exists, this option can be used to specify which axis can be "onZero" to.
	OnZeroAxisIndex int `json:"onZeroAxisIndex,omitempty"`

	// Symbol of the two ends of the axis. It could be a string, representing the same symbol for two ends; or an array
	// with two string elements, representing the two ends separately. It's set to be 'none' by default, meaning no
	// arrow for either end. If it is set to be 'arrow', there shall be two arrows. If there should only one arrow
	// at the end, it should set to be ['none', 'arrow'].
	Symbol string `json:"symbol,omitempty"`

	// Size of the arrows at two ends. The first is the width perpendicular to the axis, the next is the width parallel to the axis.
	SymbolSize []float64 `json:"symbolSize,omitempty"`

	// Arrow offset of axis. If is array, the first number is the offset of the arrow at the beginning, and the second
	// number is the offset of the arrow at the end. If is number, it means the arrows have the same offset.
	SymbolOffset []float64 `json:"symbolOffset,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}
