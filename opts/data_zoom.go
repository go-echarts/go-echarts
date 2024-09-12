package opts

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

	// Specify whether the layout of dataZoom component is horizontal or vertical. What's more, it indicates whether the horizontal axis or vertical axis is controlled by default in catesian coordinate system.
	//
	// Valid values:
	// 'horizontal': horizontal.
	// 'vertical': vertical.
	Orient string `json:"orient,omitempty"`

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

	// LabelFormatter is the formatter tool for the label.
	//
	// If it is a string, it will be a template. For instance, aaaa{value}bbbb, where {value} will be replaced by the value of actual data value.
	// It can also be a callback function. For example:
	//
	// /** @param {*} value If axis.type is 'category', `value` is the index of axis.data.
	//  *                   else `value` is current value.
	//  * @param {string} valueStr Inner formatted string.
	//  * @return {string} Returns the label formatted.
	//  labelFormatter: function (value, valueStr) {
	//     return 'aaa' + value + 'bbb';
	// }
	LabelFormatter string `json:"labelFormatter,omitempty"`

	// FilterMode Generally dataZoom component zoom or roam coordinate system
	// https://echarts.apache.org/en/option.html#dataZoom-inside.filterMode
	// through data filtering and set the windows of axes internally.
	// Possible values:
	//'filter': data that outside the window will be filtered, which may lead to some changes of windows of other axes. For each data item, it will be filtered if one of the relevant dimensions is out of the window.
	//'weakFilter': data that outside the window will be filtered, which may lead to some changes of windows of other axes. For each data item, it will be filtered only if all of the relevant dimensions are out of the same side of the window.
	//'empty': data that outside the window will be set to NaN, which will not lead to changes of windows of other axes.
	//'none': Do not filter data.
	FilterMode string `json:"filterMode,omitempty"`
}
