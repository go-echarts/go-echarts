package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Tooltip is the option set for a tooltip component.
// https://echarts.apache.org/en/option.html#tooltip
type Tooltip struct {
	// Whether to show the tooltip component, including tooltip floating layer and axisPointer.
	Show types.Bool `json:"show,omitempty"`

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
	Enterable types.Bool `json:"enterable,omitempty"`

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
	Formatter types.FuncStr `json:"formatter,omitempty"`

	// ValueFormatter Callback function for formatting the value section in tooltip.
	// valueFormatter: (value) => '$' + value.toFixed(2)
	ValueFormatter types.FuncStr `json:"valueFormatter,omitempty"`

	// The content formatter of tooltip's floating layer which supports string template and callback function.
	// See https://echarts.apache.org/en/option.html#grid.tooltip.position
	// May be a string ("inside", "top", "bottom", "left", "right") or a function of form:
	//   (point: Array, params: Object|Array.<Object>, dom: HTMLDomElement, rect: Object, size: Object) => Array
	Position types.FuncStr `json:"position,omitempty"`

	// The border color of tooltip's floating layer.
	BorderColor string `json:"borderColor,omitempty"`

	// The background color of tooltip's floating layer. e.g. 'rgba(50,50,50,0.7)'
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Configuration item for axisPointer
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}
