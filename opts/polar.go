package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Polar coordinate can be used in scatter and line chart. Every polar coordinate has an angleAxis and a radiusAxis.
type Polar struct {
	ID string `json:"id,omitempty"`

	// Zlevel value of all graphical elements in .
	// Zlevel is used to make layers with Canvas.
	// Graphical elements with different zlevel values will be placed in different Canvases, which is a common optimization technique.
	// Canvases with bigger zlevel will be placed on Canvases with smaller zlevel.

	Zlevel int `json:"zlevel,omitempty"`

	// Z value of all graphical elements in , which controls order of drawing graphical components.
	// Components with smaller z values may be overwritten by those with larger z values.
	// Z has a lower priority to Zlevel, and will not create new Canvas.
	Z int `json:"z,omitempty"`

	// Center position of Polar coordinate, the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width, and the second item to the height.
	// Example:
	// Set to absolute pixel values
	//   center: [400, 300]
	// Set to relative percent
	//   center: ['50%', '50%']
	Center []string `json:"center,omitempty"`

	// Radius of Polar coordinate. Value can be:
	Radius []string `json:"radius,omitempty"`

	// Tooltip settings in the coordinate system component.
	Tooltip Tooltip `json:"tooltip,omitempty"`
}

type PolarAxisBase struct {
	ID string `json:"id,omitempty"`
	// PolarIndex Index of radial axis in polar coordinate. It's the first axis by default.
	PolarIndex int `json:"polarIndex,omitempty"`
	// Type of axis.
	// Option:
	// 'value' Numerical axis, suitable for continuous data.
	// 'category' Category axis, suitable for discrete category data.
	// 'time' Time axis, suitable for continuous time series data.
	// 'log' Log axis, suitable for log data.
	Type        string     `json:"type,omitempty"`
	StartAngle  float64    `json:"startAngle,omitempty"`
	BoundaryGap types.Bool `json:"boundaryGap,omitempty"`
	Min         float64    `json:"min,omitempty"`
	Max         float64    `json:"max,omitempty"`
	Scale       types.Bool `json:"scale,omitempty"`
	SplitNumber int        `json:"splitNumber,omitempty"`
	MinInterval float64    `json:"minInterval,omitempty"`
	MaxInterval float64    `json:"maxInterval,omitempty"`
	Interval    float64    `json:"interval,omitempty"`
	// LogBase Base of logarithm, which is valid only for numeric axes with type: 'log'.
	LogBase float64 `json:"logBase,omitempty"`
	// Silent Set this to true, to prevent interaction with the axis.
	Silent types.Bool `json:"silent,omitempty"`
	// TriggerEvent Set this to true to enable triggering events.
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
}
