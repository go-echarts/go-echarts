package components

import "github.com/go-echarts/go-echarts/v2/primitive"

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

func (yAxis3D YAxis3D) New() *YAxis3D {
	return &YAxis3D{}
}

// YAxis is the option set for Y axis.
// https://echarts.apache.org/en/option.html#yAxis
type YAxis struct {
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

	// The index of grid which the Y axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex primitive.Int `json:"gridIndex,omitempty"`

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

func (yAxis YAxis) New() *YAxis {
	return &YAxis{
		Type: "value",
	}
}
