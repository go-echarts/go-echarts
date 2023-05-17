package opt

import "github.com/go-echarts/go-echarts/v2/primitive"

// XAxis is the option set for X axis.
// https://echarts.apache.org/en/option.html#xAxis
type XAxis struct {
	// Name of axis.
	Name primitive.String `json:"name,omitempty"`

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
	Type primitive.String `json:"type,omitempty"`

	// Set this to false to prevent the axis from showing.
	Show primitive.Bool `json:"show,omitempty"`

	// Category data, available in type: 'category' axis.
	Data primitive.Mixed `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber primitive.Int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale primitive.Bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min primitive.Mixed `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max primitive.Mixed `json:"max,omitempty"`

	// Minimum gap between split lines. For 'time' axis, MinInterval is in unit of milliseconds.
	MinInterval primitive.Float `json:"minInterval,omitempty"`

	// Maximum gap between split lines. For 'time' axis, MaxInterval is in unit of milliseconds.
	MaxInterval primitive.Float `json:"maxInterval,omitempty"`

	// The index of grid which the x axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex primitive.Int `json:"gridIndex,omitempty"`

	// Split area of X axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Settings related to axis tick.
	AxisTick *AxisTick `json:"axisTick,omitempty"`

	// Settings related to axis pointer.
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}

func (xAxis XAxis) New() *XAxis {
	return &XAxis{
		Type: "category",
	}
}
