package opts

import "github.com/go-echarts/go-echarts/v2/types"

// YAxis is the option set for Y axis.
// https://echarts.apache.org/en/option.html#yAxis
type YAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

	// AlignTicks turned on to automatically align ticks when multiple numeric y axes.
	// Only available for axes of type 'value' and 'log'.
	AlignTicks types.Bool `json:"alignTicks,omitempty"`

	// Position the position of y-axis.
	//options:
	//'left' (default)
	//'right'
	Position string `json:"position,omitempty"`

	// Location of axis name.
	//
	// Options:
	// 'start'
	// 'middle' or 'center'
	// 'end'
	NameLocation string `json:"nameLocation,omitempty"`

	// Gap between axis name and axis line.
	NameGap int `json:"nameGap,omitempty"`

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
	Show types.Bool `json:"show,omitempty"`

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

	// Minimum gap between split lines.
	//
	// For example, it can be set to be 1 to make sure axis label is show as integer.
	// {
	//    minInterval: 1
	// }
	// It is available only for axis of type 'value' or 'time'.
	MinInterval float64 `json:"minInterval,omitempty"`

	// Maximum gap between split lines.
	// For example, in time axis (type is 'time'), it can be set to be 3600 * 24 * 1000 to make sure that the gap between axis labels is less than or equal to one day.
	//{
	//    maxInterval: 3600 * 1000 * 24
	//}
	// It is available only for axis of type 'value' or 'time'.
	MaxInterval float64 `json:"maxInterval,omitempty"`

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

	// Settings related to axis pointer.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}
