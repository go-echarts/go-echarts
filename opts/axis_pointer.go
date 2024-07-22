package opts

import "github.com/go-echarts/go-echarts/v2/types"

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
	Snap types.Bool `json:"snap,omitempty"`

	Link []AxisPointerLink `json:"link,omitempty"`

	Axis string `json:"axis,omitempty"`

	Show types.Bool `json:"show,omitempty"`

	Label *Label `json:"label,omitempty"`
}

type AxisPointerLink struct {
	XAxisIndex []int  `json:"xAxisIndex,omitempty"`
	YAxisIndex []int  `json:"yAxisIndex,omitempty"`
	XAxisName  string `json:"xAxisName,omitempty"`
	YAxisName  string `json:"yAxisName,omitempty"`
}
