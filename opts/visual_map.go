package opts

import "github.com/go-echarts/go-echarts/v2/types"

// VisualMap is a type of component for visual encoding, which maps the data to visual channels.
// https://echarts.apache.org/en/option.html#visualMap
type VisualMap struct {
	// Mapping type.
	// Options: "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`

	// Whether show handles, which can be dragged to adjust "selected range".
	Calculable types.Bool `json:"calculable"`

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
	Show types.Bool `json:"show,omitempty"`

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

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Text style
	TextStyle *TextStyle `json:"textStyle,omitempty"`
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

// Piece Used to customize how to slice continuous data, and some specific view style for some pieces.
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
