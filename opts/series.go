package opts

import "fmt"

// Label contains options for a label text.
// https://echarts.apache.org/en/option.html#series-line.label
type Label struct {
	// Whether to show label.
	Show bool `json:"show"`

	// Color is the text color.
	// If set as "auto", the color will assigned as visual color, such as series color.
	Color string `json:"color,omitempty"`

	// Label position. Followings are the options:
	//
	// [x, y]
	// Use relative percentage, or absolute pixel values to represent position of label
	// relative to top-left corner of bounding box. For example:
	//
	// Absolute pixel values: position: [10, 10],
	// Relative percentage: position: ["50%", "50%"]
	//
	// "top"
	// "left"
	// "right"
	// "bottom"
	// "inside"
	// "insideLeft"
	// "insideRight"
	// "insideTop"
	// "insideBottom"
	// "insideTopLeft"
	// "insideBottomLeft"
	// "insideTopRight"
	// "insideBottomRight"
	Position string `json:"position,omitempty"`

	// Data label formatter, which supports string template and callback function.
	// In either form, \n is supported to represent a new line.
	// String template, Model variation includes:
	//
	// {a}: series name.
	// {b}: the name of a data item.
	// {c}: the value of a data item.
	// {@xxx}: the value of a dimension named"xxx", for example,{@product}refers the value of"product"` dimension.
	// {@[n]}: the value of a dimension at the index ofn, for example,{@[3]}` refers the value at dimensions[3].
	Formatter string `json:"formatter,omitempty"`
}

// LabelLine Configuration of label guide line.
type LabelLine struct {
	// Whether to show the label guide line.
	Show bool `json:"show"`
	// Whether to show the label guide line above the corresponding element.
	ShowAbove bool `json:"showAbove"`
	// The length of the second segment of guide line.
	Length2 float64 `json:"length2,omitempty"`
	// smoothness of guide line.
	Smooth bool `json:"smooth"`
	// Minimum turn angle between two segments of guide line
	MinTurnAngle float64 `json:"minTurnAngle,omitempty"`
	// The style of label line
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// Emphasis is the style when it is highlighted, like being hovered by mouse, or highlighted via legend connect.
type Emphasis struct {
	// the emphasis style of label
	Label *Label `json:"label,omitempty"`

	// the emphasis style of item
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// ItemStyle
type ItemStyle struct {
	// Color of chart
	// Kline Up candle color
	Color string `json:"color,omitempty"`

	// Kline Down candle color
	Color0 string `json:"color0,omitempty"`

	// BorderColor is the hart border color
	// Kline  Up candle border color
	BorderColor string `json:"borderColor,omitempty"`

	// Kline Down candle border color
	BorderColor0 string `json:"borderColor0,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`
}

// MarkLines represents a series of marklines.
type MarkLines struct {
	Data []interface{} `json:"data,omitempty"`
	MarkLineStyle
}

// MarkLineStyle contains styling options for a MarkLine.
type MarkLineStyle struct {
	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol []string `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`

	// Mark line text options.
	Label *Label `json:"label,omitempty"`
}

// MarkLineNameTypeItem represents type for a MarkLine.
type MarkLineNameTypeItem struct {
	// Mark line name.
	Name string `json:"name,omitempty"`

	// Mark line type, options: "average", "min", "max".
	Type string `json:"type,omitempty"`

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLineNameYAxisItem
type MarkLineNameYAxisItem struct {
	// Mark line name
	Name string `json:"name,omitempty"`

	// Y axis data
	YAxis interface{} `json:"yAxis,omitempty"`

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLineNameXAxisItem
type MarkLineNameXAxisItem struct {
	// Mark line name
	Name string `json:"name,omitempty"`

	// X axis data
	XAxis interface{} `json:"xAxis,omitempty"`

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLineNameCoordItem represents coordinates for a MarkLine.
type MarkLineNameCoordItem struct {
	// Mark line name
	Name string `json:"name,omitempty"`

	// Mark line start coordinate
	Coordinate0 []interface{}

	// Mark line end coordinate
	Coordinate1 []interface{}

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkPoints represents a series of markpoints.
type MarkPoints struct {
	Data []interface{} `json:"data,omitempty"`
	MarkPointStyle
}

// MarkPointStyle contains styling options for a MarkPoint.
type MarkPointStyle struct {
	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol []string `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`

	// Mark point text options.
	Label *Label `json:"label,omitempty"`
}

// MarkPointNameTypeItem represents type for a MarkPoint.
type MarkPointNameTypeItem struct {
	// Name of markpoint
	Name string `json:"name,omitempty"`

	// Mark point type, options: "average", "min", "max".
	Type string `json:"type,omitempty"`

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkPointNameCoordItem represents coordinates for a MarkPoint.
type MarkPointNameCoordItem struct {
	// Name of markpoint
	Name string `json:"name,omitempty"`

	// Mark point coordinate
	Coordinate []interface{} `json:"coord,omitempty"`

	// Works only when type is assigned.
	// It is used to state the dimension used to calculate maximum value or minimum value.
	// It may be the direct name of a dimension, like x,
	// or angle for line charts, or open, or close for candlestick charts.
	ValueDim string `json:"valueDim,omitempty"`

	// Mark point text options.
	Label *Label `json:"label,omitempty"`
}

// RippleEffect is the option set for the ripple effect.
type RippleEffect struct {
	// The period duration of animation, in seconds.
	// default 4(s)
	Period float32 `json:"period,omitempty"`

	// The maximum zooming scale of ripples in animation.
	// default 2.5
	Scale float32 `json:"scale,omitempty"`

	// The brush type for ripples. options: "stroke" and "fill".
	// default "fill"
	BrushType string `json:"brushType,omitempty"`
}

// LineStyle is the option set for a link style component.
type LineStyle struct {
	// Line color
	Color string `json:"color,omitempty"`

	// Width of line. default 1
	Width float32 `json:"width,omitempty"`

	// Type of line，options: "solid", "dashed", "dotted". default "solid"
	Type string `json:"type,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`

	// Curveness of edge. The values from 0 to 1 could be set.
	// it would be larger as the the value becomes larger. default 0
	Curveness float32 `json:"curveness,omitempty"`
}

// AreaStyle is the option set for an area style component.
type AreaStyle struct {
	// Fill area color.
	Color string `json:"color,omitempty"`

	// Opacity of the component. Supports value from 0 to 1, and the component will not be drawn when set to 0.
	Opacity float32 `json:"opacity,omitempty"`
}

// Configuration items about force-directed layout. Force-directed layout simulates
// spring/charge model, which will add a repulsion between 2 nodes and add a attraction
// between 2 nodes of each edge. In each iteration nodes will move under the effect
// of repulsion and attraction. After several iterations, the nodes will be static in a
// balanced position. As a result, the energy local minimum of this whole model will be realized.
// The result of force-directed layout has a good symmetries and clustering, which is also aesthetically pleasing.
type GraphForce struct {
	// The initial layout before force-directed layout, which will influence on the result of force-directed layout.
	// It defaults not to do any layout and use x, y provided in node as the position of node.
	// If it doesn"t exist, the position will be generated randomly.
	// You can also use circular layout "circular".
	InitLayout string `json:"initLayout,omitempty"`

	// The repulsion factor between nodes. The repulsion will be stronger and the distance
	// between 2 nodes becomes further as this value becomes larger.
	// It can be an array to represent the range of repulsion. In this case larger value have larger
	// repulsion and smaller value will have smaller repulsion.
	Repulsion float32 `json:"repulsion,omitempty"`

	// The gravity factor enforcing nodes approach to the center. The nodes will be
	// closer to the center as the value becomes larger. default 0.1
	Gravity float32 `json:"gravity,omitempty"`

	// The distance between 2 nodes on edge. This distance is also affected by repulsion.
	// It can be an array to represent the range of edge length. In this case edge with larger
	// value will be shorter, which means two nodes are closer. And edge with smaller value will be longer.
	// default 30
	EdgeLength float32 `json:"edgeLength,omitempty"`
}

// Leaf node special configuration, the leaf node and non-leaf node label location is different.
type TreeLeaves struct {
	// The style setting of the text label in a single bar.
	Label *Label `json:"label,omitempty"`

	// LineStyle settings in this series data.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Emphasis settings in this series data.
	Emphasis *Emphasis `json:"emphasis,omitempty"`
}

// RGBColor returns the color with RGB format
func RGBColor(r, g, b uint16) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}

// RGBAColor returns the color with RGBA format
func RGBAColor(r, g, b uint16, a float32) string {
	return fmt.Sprintf("rgba(%d,%d,%d,%f)", r, g, b, a)
}

// HSLColor returns the color with HSL format
func HSLColor(h, s, l float32) string {
	return fmt.Sprintf("hsl(%f,%f%%,%f%%)", h, s, l)
}

// HSLAColor returns the color with HSLA format
func HSLAColor(h, s, l, a float32) string {
	return fmt.Sprintf("hsla(%f,%f%%,%f%%,%f)", h, s, l, a)
}
